package route

import (
	networkingv1beta1 "k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

var DefaultAnno map[string]string

func initDefaultAnnotation() {
	DefaultAnno = make(map[string]string)
	DefaultAnno["nginx.ingress.kubernetes.io/rewrite-target"] = "/"
}

// helper function to quickly build metav1.ObjectMeta objcet
func BuildObjectMeta(ns string, name string, anno map[string]string) metav1.ObjectMeta {
	// merge the default annotations and the inputed annotations
	mergedAnno := make(map[string]string)
	for key, val := range DefaultAnno {
		mergedAnno[key] = val
	}
	for key, val := range anno {
		mergedAnno[key] = val
	}
	// build a metav1.ObjectMeta object
	objMeta := metav1.ObjectMeta{
		Name:        name,
		Namespace:   ns,
		Annotations: mergedAnno,
	}
	return objMeta
}

// helper function to quickly build networkingv1beta1.IngressSpec object
func BuildIngressSpec(path string, wfName string) networkingv1beta1.IngressSpec {
	// fix path var if the head is not '/'
	if path[0] != '/' {
		path = "/" + path
	}
	// used for object initialtion
	pathType := networkingv1beta1.PathType("Prefix")
	// build a networkingv1beta1.IngressSpec object
	IngSpec := networkingv1beta1.IngressSpec{
		Rules: []networkingv1beta1.IngressRule{
			{
				IngressRuleValue: networkingv1beta1.IngressRuleValue{
					HTTP: &networkingv1beta1.HTTPIngressRuleValue{
						Paths: []networkingv1beta1.HTTPIngressPath{
							{
								Path:     path,
								PathType: &pathType,
								Backend: networkingv1beta1.IngressBackend{
									ServiceName: wfName,
									ServicePort: intstr.Parse("80"),
								},
							},
						},
					},
				},
			},
		},
	}
	return IngSpec
}

// helper function to quickly build IngressName
func BuildIngressName(wfName string) string {
	return "tass-ingress-" + wfName
}

func init() {
	initDefaultAnnotation()
}
