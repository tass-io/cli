package function

import (
	"context"
	"errors"
	"fmt"

	"github.com/tass-io/cli/pkg/storagesvc"
	serverlessv1alpha1 "github.com/tass-io/tass-operator/api/v1alpha1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	runtimeClient "sigs.k8s.io/controller-runtime/pkg/client"
)

type GetFunction struct {
	name string
	ns   string
	fn   *serverlessv1alpha1.Function
}

// do is the business logic of creating a Function
func (gf *GetFunction) do() error {
	err := gf.get()
	if k8serrors.IsNotFound(err) {
		return errors.New("A function with the name not existed")
	}
	if runtimeClient.IgnoreNotFound(err) != nil {
		// Get Function failed
		return err
	}
	return gf.print()
}

// get gets the Function by name and namespace
func (cf *GetFunction) get() error {
	err := client.Get(context.Background(), runtimeClient.ObjectKey{
		Namespace: cf.ns,
		Name:      cf.name,
	}, cf.fn)
	return err
}

// print prints the information about function code
func (gf *GetFunction) print() error {
	code, err := gf.getCode()
	if err != nil {
		return err
	}
	fmt.Println("The code of the function:")
	fmt.Println(code)
	return nil
}

// getCode gets the function code from the storage service
func (gf *GetFunction) getCode() (string, error) {
	return storagesvc.Get(gf.ns, gf.name)
}
