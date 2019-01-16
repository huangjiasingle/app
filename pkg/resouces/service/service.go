package service

import (
	appv1 "app/pkg/apis/app/v1"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func New(app *appv1.App) *corev1.Service {

	return &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      app.Name,
			Namespace: app.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(app, schema.GroupVersionKind{
					Group:   appv1.SchemeGroupVersion.Group,
					Version: appv1.SchemeGroupVersion.Version,
					Kind:    "App",
				}),
			},
		},
		Spec: corev1.ServiceSpec{
			Ports: app.Spec.Ports,
			Selector: map[string]string{
				"app.example.com/v1": app.Name,
			},
		},
	}
}
