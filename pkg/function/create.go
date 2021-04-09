package function

import (
	"context"
	"errors"

	"github.com/tass-io/cli/pkg/storagesvc"
	serverlessv1alpha1 "github.com/tass-io/tass-operator/api/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type CreateFunction struct {
	name   string
	ns     string
	code   string
	client client.Client
	fn     *serverlessv1alpha1.Function
}

// do is the business logic of creating a Function
// Note that we need to first store the function code then complete the CRD creation
// The commit point is the CRD Creation
func (cf *CreateFunction) do() error {
	err := cf.get()
	if client.IgnoreNotFound(err) != nil {
		// Get Function failed
		return err
	}
	if cf.name == cf.fn.Name {
		return errors.New("A function with the same name already exists")
	}
	if err = cf.store(); err != nil {
		// Store the function code failed
		return err
	}
	return cf.complete()
}

// get gets the Function by name and namespace
func (cf *CreateFunction) get() error {
	err := cf.client.Get(context.Background(), client.ObjectKey{
		Namespace: cf.ns,
		Name:      cf.name,
	}, cf.fn)
	return err
}

// store stores the source code of the function
func (cf *CreateFunction) store() error {
	return storagesvc.Set(cf.ns, cf.name, cf.code)
}

// complete creates a Function, business logic should be done before calling this function
func (cf *CreateFunction) complete() error {
	cf.fn = &serverlessv1alpha1.Function{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: cf.ns,
			Name:      cf.name,
		},
		// FIXME: Update this field later
		Spec: serverlessv1alpha1.FunctionSpec{
			Environment: "Golang",
		},
	}
	return cf.client.Create(context.Background(), cf.fn)
}
