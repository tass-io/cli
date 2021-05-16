package client

import (
	"os"

	"github.com/tass-io/cli/pkg/logging"
	"k8s.io/apimachinery/pkg/runtime"

	serverlessv1alpha1 "github.com/tass-io/tass-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	networkingv1beta1 "k8s.io/api/networking/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

var tassClient *client.Client
var netClient *client.Client
var log = logging.Log

// GetCRDClient returns the CRD client pointer
func GetCRDClient() *client.Client {
	return tassClient
}

// GetNetClient returns the Net client pointer
func GetNetClinet() *client.Client {
	return netClient
}

// initCRDClient create a client to perform CRUD operations as well as default Kind on a Kubernetes cluster.
// In order to call the recognize CRD types,
// a scheme that has custom operator types registered for the Client is set.
func initCRDClient() {
	scheme := runtime.NewScheme()
	if err := serverlessv1alpha1.AddToScheme(scheme); err != nil {
		log.Error(err)
		os.Exit(1)
	}
	if err := corev1.AddToScheme(scheme); err != nil {
		log.Error(err)
		os.Exit(1)
	}
	cl, err := client.New(config.GetConfigOrDie(), client.Options{Scheme: scheme})
	if err != nil {
		log.Error("failed to create client")
		os.Exit(1)
	}
	tassClient = &cl
}

// initNetClient create a client to perform CRUD operations as well as default Kind on a Kubernetes cluster.
// In order to call the recognize Net types,
// a scheme that corresponding to `networking` types registered for the Client is set.
func initNetClient() {
	scheme := runtime.NewScheme()
	if err := networkingv1beta1.AddToScheme(scheme); err != nil {
		log.Error(err)
		os.Exit(1)
	}
	cl, err := client.New(config.GetConfigOrDie(), client.Options{Scheme: scheme})
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	netClient = &cl
}

func init() {
	initCRDClient()
	initNetClient()
}
