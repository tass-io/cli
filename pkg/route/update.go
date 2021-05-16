package route

import (
	"context"
	"errors"

	networkingv1beta1 "k8s.io/api/networking/v1beta1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"

	runtimeClient "sigs.k8s.io/controller-runtime/pkg/client"
)

type UpdateRoute struct {
	wfn  string
	ns   string
	path string
	ig   *networkingv1beta1.Ingress
}

// do is the business logic of updating a Route
func (ug *UpdateRoute) do() error {
	err := ug.get()
	if k8serrors.IsNotFound(err) {
		return errors.New("a route with the name not existed")
	}
	if runtimeClient.IgnoreNotFound(err) != nil {
		// Get Route failed
		return err
	}
	ug.ig.Spec.Rules[0].HTTP.Paths[0].Path = ug.path
	return ug.complete()
}

// get gets the Route by name and namespace
func (ug *UpdateRoute) get() error {
	err := client.Get(context.Background(), runtimeClient.ObjectKey{
		Namespace: ug.ns,
		Name:      BuildIngressName(ug.wfn),
	}, ug.ig)
	return err
}

// complete updates a Route, business logic should be done before calling this route
func (ug *UpdateRoute) complete() error {
	return client.Update(context.Background(), ug.ig)
}
