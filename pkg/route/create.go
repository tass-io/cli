package route

import (
	"context"
	"errors"

	networkingv1beta1 "k8s.io/api/networking/v1beta1"

	runtimeClient "sigs.k8s.io/controller-runtime/pkg/client"
)

type CreateRoute struct {
	wfn  string
	ns   string
	path string
	ig   *networkingv1beta1.Ingress
}

// do is the business logic of creating a Route
// The commit point is the Ingress Creation
func (cr *CreateRoute) do() error {
	err := cr.get()
	if runtimeClient.IgnoreNotFound(err) != nil {
		// Get Route failed
		return err
	}
	if BuildIngressName(cr.wfn) == cr.ig.Name {
		return errors.New("a route with the same name already exists")
	}
	return cr.complete()
}

// get gets the Ingress by name and namespace
func (cr *CreateRoute) get() error {
	err := client.Get(context.Background(), runtimeClient.ObjectKey{
		Namespace: cr.ns,
		Name:      BuildIngressName(cr.wfn),
	}, cr.ig)
	return err
}

// complete creates a Route, business logic should be done before calling this Route
func (cr *CreateRoute) complete() error {
	cr.ig = &networkingv1beta1.Ingress{
		ObjectMeta: BuildObjectMeta(cr.ns, BuildIngressName(cr.wfn), nil),
		// FIXME: Update this field later
		Spec: BuildIngressSpec(cr.path, cr.wfn),
	}
	return client.Create(context.Background(), cr.ig)
}
