package route

import (
	"context"

	networkingv1beta1 "k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeleteRoute struct {
	wfn string
	ns  string
}

// do is the business logic of deleting a Route
func (dr *DeleteRoute) do() error {
	return dr.complete()
}

// complete deletes a Route
func (dr *DeleteRoute) complete() error {
	ig := &networkingv1beta1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: dr.ns,
			Name:      BuildIngressName(dr.wfn),
		},
	}
	return client.Delete(context.Background(), ig)
}
