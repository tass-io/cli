package function

import (
	"context"
	"errors"

	"github.com/tass-io/cli/pkg/storagesvc"
	serverlessv1alpha1 "github.com/tass-io/tass-operator/api/v1alpha1"

	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtimeClient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/tass-io/cli/pkg/tools/base64"
)

type CreateFunction struct {
	name string
	ns   string
	code string
	fn   *serverlessv1alpha1.Function
	mock bool
}

// do is the business logic of creating a Function
// Note that we need to first store the function code then complete the CRD creation
// The commit point is the CRD Creation
func (cf *CreateFunction) do() error {
	err := cf.get()
	if runtimeClient.IgnoreNotFound(err) != nil {
		// Get Function failed
		return err
	}
	if cf.name == cf.fn.Name {
		return errors.New("a function with the same name already exists")
	}
	if err = cf.store(); err != nil {
		// Store the function code failed
		return err
	}
	return cf.complete()
}

// get gets the Function by name and namespace
func (cf *CreateFunction) get() error {
	err := client.Get(context.Background(), runtimeClient.ObjectKey{
		Namespace: cf.ns,
		Name:      cf.name,
	}, cf.fn)
	return err
}

// store stores the source code of the function
func (cf *CreateFunction) store() error {
	if cf.mock {
		return cf.mockStore()
	}
	return storagesvc.Set(cf.ns, cf.name, cf.code)
}

// mockStore stores a prepared zipped code of the function
func (cf *CreateFunction) mockStore() error {
	fileName := "mock/plugin-golang-wrapper.zip"
	mockedCode, err := base64.EncodeUserCode(fileName)
	if err != nil {
		return err
	}
	return storagesvc.CodeSet(cf.ns, cf.name, mockedCode)
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
			Resource: serverlessv1alpha1.Resource{
				// FIXME: Update this field later
				ResourceCPU:    resource.Quantity{},
				ResourceMemory: resource.Quantity{},
			},
		},
	}
	return client.Create(context.Background(), cf.fn)
}
