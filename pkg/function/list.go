package function

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	serverlessv1alpha1 "github.com/tass-io/tass-operator/api/v1alpha1"
	runtimeClient "sigs.k8s.io/controller-runtime/pkg/client"
)

type ListFunctions struct {
	ns     string
	fnList *serverlessv1alpha1.FunctionList
}

// do is the business logic of creating a Function
func (lf *ListFunctions) do() error {
	err := lf.complete()
	if runtimeClient.IgnoreNotFound(err) != nil {
		// Get Function failed
		return err
	}
	lf.print()
	return nil
}

// complete gets the Function by name and namespace
func (lf *ListFunctions) complete() error {
	err := client.List(context.Background(), lf.fnList, &runtimeClient.ListOptions{Namespace: lf.ns})
	return err
}

// print prints the information about commands
func (lf *ListFunctions) print() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	defer w.Flush()
	fmt.Fprintf(w, "%v\t%v\t%v\n", "NAMESPACE", "NAME", "ENV")
	for _, fn := range lf.fnList.Items {
		fmt.Fprintf(w, "%v\t%v\t%v\n", fn.ObjectMeta.Namespace, fn.ObjectMeta.Name, fn.Spec.Environment)
	}
}
