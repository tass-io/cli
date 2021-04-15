package function

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"net/http"

	serverlessv1alpha1 "github.com/tass-io/tass-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtimeClient "sigs.k8s.io/controller-runtime/pkg/client"
)

type TestFunction struct {
	name string
	ns   string
	fn   *serverlessv1alpha1.Function
	svc  *corev1.Service
}

// do is the business logic of testing a Function
// In order to test a Function, there are some steps to do
// 1. Create a temporary Pod and its corresponding Service for test
// 2. Test the health of the Pod
// 3. Do the test (TODO: Start the image and run the function)
// 4. Clean the temporary resource
func (tf *TestFunction) do() error {
	// 1. Create a temporary Pod and its corresponding Service for test
	if err := tf.create(); err != nil {
		return err
	}
	// 2. Test the health of the Pod with 120 second timeout
	if err := tf.checkPodStatus(120); err != nil {
		log.Error("The Pod is not ready.")
		return err
	}
	// 3. Do the test (TODO: Start the image and run the function)
	if err := tf.request(); err != nil {
		return err
	}
	// 4. Clean the temporary resource
	if err := tf.clean(); err != nil {
		return err
	}
	return nil
}

// create creates resources of a Service and a Pod
// These resource are temporary, they will be deleted when the test finishes
func (tf *TestFunction) create() error {
	tf.name += "-test"
	labels := map[string]string{
		"kind":     "test",
		"function": tf.name,
	}

	svc := createSvcConf(labels, tf.name, tf.ns)
	pod := createPoConf(labels, tf.name, tf.ns)

	log.Info("Creating a test Service...")
	if err := client.Create(context.Background(), svc); err != nil {
		log.Error("Create Service failed")
		return err
	}

	log.Info("Creating a test Pod...")
	if err := client.Create(context.Background(), pod); err != nil {
		log.Error("Create Pod failed")
		return err
	}
	return nil
}

// createPoConf creates a test Pod config for testing the Function
func createPoConf(labels map[string]string, name, ns string) *corev1.Pod {
	return newPod(labels, name, ns)
}

// createSvcConf creates a test Service config for testing the Function
func createSvcConf(labels map[string]string, name, ns string) *corev1.Service {
	return newSvc(labels, name, ns, int32(80))
}

// checkPodStatus waits for the pod status being ready until timeout(second)
func (tf *TestFunction) checkPodStatus(timeout int) error {
	for i := 0; i < timeout; i += 5 {
		log.Info("Waiting for Pod creation...")
		isReady := isPodReady(tf.name, tf.ns)
		if isReady {
			log.Info("The Pod is ready.")
			return nil
		}
		time.Sleep(time.Second * 5)
	}
	return errors.New("Waiting for Pod to be ready timeout, please check the Pod status in cluster")
}

// request sends a request and test the function. The pod is
// found in the specified namespace by labelSelector. The pod's port
// is found by looking for a service in the same namespace and using
// its targetPort. Once the port forward is started, wait for it to
// start accepting connections before returning.
func (tf *TestFunction) request() error {
	// get test function Service resource
	err := tf.getSvc()
	if k8serrors.IsNotFound(err) {
		return errors.New("The Service of the function is not found.")
	}
	if runtimeClient.IgnoreNotFound(err) != nil {
		return err
	}
	ip := tf.svc.Spec.ClusterIP
	// there should be only one port
	var targetPort string
	for _, servicePort := range tf.svc.Spec.Ports {
		targetPort = servicePort.TargetPort.String()
	}

	log.Info("Test the Function.")

	// Send http request
	// TODO: This is a test address
	log.Info("The Function response is:")
	address := "http://" + ip + ":" + targetPort + "/get"
	resp, err := http.Get(address)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// TODO: Return the result directly
	fmt.Println(string(body))
	return nil
}

// getSvc gets the Service by name and namespace
func (tf *TestFunction) getSvc() error {
	err := client.Get(context.Background(), runtimeClient.ObjectKey{
		Namespace: tf.ns,
		Name:      tf.name,
	}, tf.svc)
	return err
}

// clean deletes the test resource of Service and Pod
// If the clean action founds the resource not found, it will ignore the error
func (tf *TestFunction) clean() error {
	log.Info("Clean the test resource...")
	po := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: tf.ns,
			Name:      tf.name,
		},
	}
	if err := client.Delete(context.Background(), po); runtimeClient.IgnoreNotFound(err) != nil {
		return err
	}
	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: tf.ns,
			Name:      tf.name,
		},
	}
	if err := client.Delete(context.Background(), svc); runtimeClient.IgnoreNotFound(err) != nil {
		return err
	}
	return nil
}
