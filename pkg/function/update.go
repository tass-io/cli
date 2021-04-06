package function

import (
	"context"
	"errors"

	serverlessv1alpha1 "github.com/tass-io/tass-operator/api/v1alpha1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

type UpdateFunction struct {
	name   string
	ns     string
	code   string
	client client.Client
	fn     *serverlessv1alpha1.Function
}

// do is the business logic of updating a Function
func (uf *UpdateFunction) do() error {
	err := uf.get()
	if k8serrors.IsNotFound(err) {
		return errors.New("A function with the name not existed")
	}
	if client.IgnoreNotFound(err) != nil {
		// Get Function failed
		return err
	}
	return uf.complete()
}

// get gets the Function by name and namespace
func (uf *UpdateFunction) get() error {
	err := uf.client.Get(context.Background(), client.ObjectKey{
		Namespace: uf.ns,
		Name:      uf.name,
	}, uf.fn)
	return err
}

// complete updates a Function, business logic should be done before calling this function
func (uf *UpdateFunction) complete() error {
	// FIXME: MOCK a update logic, update this logic later
	uf.fn.Spec.Domain += "-update"
	return uf.client.Update(context.Background(), uf.fn)
}
