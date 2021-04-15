package function

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	runtimeClient "sigs.k8s.io/controller-runtime/pkg/client"
)

// newPod returns the config of creating a new Pod
func newPod(labels map[string]string, name, ns string) *corev1.Pod {
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: ns,
			Labels:    labels,
		},
		Spec: corev1.PodSpec{
			// TODO: Update the image future
			Containers: []corev1.Container{
				{
					Name:  "httpbin",
					Image: "kennethreitz/httpbin",
					Ports: []corev1.ContainerPort{{
						ContainerPort: 80,
						Protocol:      "TCP",
					}},
				},
			},
			RestartPolicy: corev1.RestartPolicyOnFailure,
		},
	}
}

// newSvc returns the config of creating a new Service
func newSvc(labels map[string]string, name, ns string, port int32) *corev1.Service {
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: ns,
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Selector: labels,
			Ports: []corev1.ServicePort{
				{
					Protocol: "TCP",
					Port:     port,
					TargetPort: intstr.IntOrString{
						Type:   0,
						IntVal: port,
					},
				},
			},
		},
	}
}

// isPodReady checks both all containers in a pod are ready and whether
// the .metadata.DeletionTimestamp is nil.
func isPodReady(name, ns string) bool {
	po, err := getPod(name, ns)
	if err != nil {
		return false
	}
	// pod is in "Terminating" status if deletionTimestamp is not nil
	// https://github.com/kubernetes/kubernetes/issues/61376
	if po.ObjectMeta.DeletionTimestamp != nil {
		return false
	}

	// pod does not have an IP address allocated to it yet
	if po.Status.PodIP == "" {
		return false
	}

	for _, cStatus := range po.Status.ContainerStatuses {
		if !cStatus.Ready {
			return false
		}
	}

	return true
}

// getPod returns the pod info
func getPod(name, ns string) (*corev1.Pod, error) {
	pod := &corev1.Pod{}
	if err := client.Get(context.Background(), runtimeClient.ObjectKey{
		Namespace: ns,
		Name:      name,
	}, pod); err != nil {
		return nil, err
	}
	return pod, nil
}
