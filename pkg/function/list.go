package function

import (
	"context"

	serverlessv1alpha1 "github.com/tass-io/tass-operator/api/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ListFunctions struct {
	ns     string
	client client.Client
	fnList *serverlessv1alpha1.FunctionList
}

// do is the business logic of creating a Function
// TODO: Get should return source code of the function
func (lf *ListFunctions) do() error {
	err := lf.complete()
	if client.IgnoreNotFound(err) != nil {
		// Get Function failed
		return err
	}
	return nil
}

// complete gets the Function by name and namespace
func (lf *ListFunctions) complete() error {
	err := lf.client.List(context.Background(), lf.fnList, &client.ListOptions{Namespace: lf.ns})
	return err
}
