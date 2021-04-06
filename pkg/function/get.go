package function

import (
	"context"
	"errors"

	serverlessv1alpha1 "github.com/tass-io/tass-operator/api/v1alpha1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type GetFunction struct {
	name   string
	ns     string
	client client.Client
	fn     *serverlessv1alpha1.Function
}

// do is the business logic of creating a Function
// TODO: Get should return source code of the function
func (gf *GetFunction) do() error {
	err := gf.get()
	if k8serrors.IsNotFound(err) {
		return errors.New("A function with the name not existed")
	}
	if client.IgnoreNotFound(err) != nil {
		// Get Function failed
		return err
	}
	return nil
}

// get gets the Function by name and namespace
func (cf *GetFunction) get() error {
	err := cf.client.Get(context.Background(), client.ObjectKey{
		Namespace: cf.ns,
		Name:      cf.name,
	}, cf.fn)
	return err
}
