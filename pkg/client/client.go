package client

import (
	"fmt"
	"os"

	"k8s.io/apimachinery/pkg/runtime"

	serverlessv1alpha1 "github.com/tass-io/tass-operator/api/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

var (
	c *client.Client
)

// GetCRDClient returns the CRD client pointer
func GetCRDClient() *client.Client {
	return c
}

// initCRDClient create a client to perform CRUD operations on a Kubernetes cluster.
// In order to call the recognize CRD types,
// a scheme that has custom operator types registered for the Client is set.
func initCRDClient() {
	scheme := runtime.NewScheme()
	err := serverlessv1alpha1.AddToScheme(scheme)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cl, err := client.New(config.GetConfigOrDie(), client.Options{Scheme: scheme})
	if err != nil {
		fmt.Println("failed to create client")
		os.Exit(1)
	}
	c = &cl
}

func init() {
	initCRDClient()
}
