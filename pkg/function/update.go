package function

import (
	"context"
	"errors"

	"github.com/tass-io/cli/pkg/storagesvc"
	serverlessv1alpha1 "github.com/tass-io/tass-operator/api/v1alpha1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"

	runtimeClient "sigs.k8s.io/controller-runtime/pkg/client"
)

type UpdateFunction struct {
	name string
	ns   string
	code string
	fn   *serverlessv1alpha1.Function
}

// do is the business logic of updating a Function
func (uf *UpdateFunction) do() error {
	err := uf.get()
	if k8serrors.IsNotFound(err) {
		return errors.New("A function with the name not existed")
	}
	if runtimeClient.IgnoreNotFound(err) != nil {
		// Get Function failed
		return err
	}
	if err := uf.store(); err != nil {
		return err
	}
	return uf.complete()
}

// get gets the Function by name and namespace
func (uf *UpdateFunction) get() error {
	err := client.Get(context.Background(), runtimeClient.ObjectKey{
		Namespace: uf.ns,
		Name:      uf.name,
	}, uf.fn)
	return err
}

// store stores the source code of the function
// store covers the old function code, regardless of the old version
func (uf *UpdateFunction) store() error {
	return storagesvc.Set(uf.ns, uf.name, uf.code)
}

// complete updates a Function, business logic should be done before calling this function
func (uf *UpdateFunction) complete() error {
	return client.Update(context.Background(), uf.fn)
}
