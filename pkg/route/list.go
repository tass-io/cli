package route

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	networkingv1beta1 "k8s.io/api/networking/v1beta1"
	runtimeClient "sigs.k8s.io/controller-runtime/pkg/client"
)

type ListRoutes struct {
	ns     string
	igList *networkingv1beta1.IngressList
}

// do is the business logic of creating a Route
func (lr *ListRoutes) do() error {
	err := lr.complete()
	if runtimeClient.IgnoreNotFound(err) != nil {
		// Get Route failed
		return err
	}
	lr.print()
	return nil
}

// complete gets the Route by name and namespace
func (lr *ListRoutes) complete() error {
	err := client.List(context.Background(), lr.igList, &runtimeClient.ListOptions{Namespace: lr.ns})
	return err
}

// print prints the information about commands
func (lr *ListRoutes) print() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	defer w.Flush()
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", "NAMESPACE", "NAME", "PATH", "ANNOTATIONS")
	for _, ig := range lr.igList.Items {
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", ig.ObjectMeta.Namespace, ig.ObjectMeta.Name, ig.Spec.Rules[0].HTTP.Paths[0].Path, len(ig.ObjectMeta.Annotations))
	}
}
