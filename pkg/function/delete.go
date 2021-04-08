package function

import (
	"context"

	"github.com/tass-io/cli/pkg/storagesvc"
	serverlessv1alpha1 "github.com/tass-io/tass-operator/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type DeleteFunction struct {
	name   string
	ns     string
	client client.Client
}

// do is the business logic of deleting a Function
func (df *DeleteFunction) do() error {
	return df.complete()
}

// complete deletes a Function
func (df *DeleteFunction) complete() error {
	fn := &serverlessv1alpha1.Function{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: df.ns,
			Name:      df.name,
		},
	}
	if err := df.clear(); err != nil {
		return err
	}
	return df.client.Delete(context.Background(), fn)
}

// clear deletes the code of the function in storage service
func (df *DeleteFunction) clear() error {
	return storagesvc.Delete(df.ns, df.name)
}
