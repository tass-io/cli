package function

import (
	"context"
	"errors"

	serverlessv1alpha1 "github.com/tass-io/tass-operator/api/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type CreateFunction struct {
	name   string
	ns     string
	domain string
	code   string
	client client.Client
	fn     *serverlessv1alpha1.Function
}

// do is the business logic of creating a Function
func (cf *CreateFunction) do() error {
	err := cf.get()
	if client.IgnoreNotFound(err) != nil {
		// Get Function failed
		return err
	}
	if cf.name == cf.fn.Name {
		return errors.New("A function with the same name already exists")
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

// complete creates a Function, business logic should be done before calling this function
func (cf *CreateFunction) complete() error {
	cf.fn = &serverlessv1alpha1.Function{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: cf.ns,
			Name:      cf.name,
		},
		// FIXME: Update this field later
		Spec: serverlessv1alpha1.FunctionSpec{
			Domain:      "cli",
			Environment: "Golang",
			Command:     "WHATS HAPPENNING",
		},
	}
	return cf.client.Create(context.Background(), cf.fn)
}
