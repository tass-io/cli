package function

import (
	"context"
	"errors"

	"github.com/tass-io/cli/pkg/agent"
	serverlessv1alpha1 "github.com/tass-io/tass-operator/api/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func doCreate() error {
	fn, err := getFn()
	if client.IgnoreNotFound(err) != nil {
		// Get Function failed
		return err
	}
	if fn.Name == fnName {
		return errors.New("A function with the same name already exists")
	}
	return createFn()
}

func getFn() (*serverlessv1alpha1.Function, error) {
	c := *agent.GetCRDClient()
	fn := &serverlessv1alpha1.Function{}
	err := c.Get(context.Background(), client.ObjectKey{
		Namespace: fnNamespace,
		Name:      fnName,
	}, fn)
	return fn, err
}

func createFn() error {
	c := *agent.GetCRDClient()
	fn := &serverlessv1alpha1.Function{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: fnNamespace,
			Name:      fnName,
		},
		// FIXME: Update this field later
		Spec: serverlessv1alpha1.FunctionSpec{
			Domain:      "cli",
			Environment: "Golang",
			Command:     "WHATS HAPPENNING",
		},
	}
	return c.Create(context.Background(), fn)
}
