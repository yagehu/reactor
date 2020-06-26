package mapper

import (
	appsv1 "k8s.io/api/apps/v1"

	"github.com/yagehu/reactor/internal/gateway/kubernetes/entity"
)

func ToKubernetesDeployment(x entity.Deployment) appsv1.Deployment {
	return appsv1.Deployment{
		TypeMeta:   ToKubernetesTypeMeta(x.TypeMeta),
		ObjectMeta: ToKubernetesObjectMeta(x.ObjectMeta),
		Spec:       ToKubernetesDeploymentSpec(x.Spec),
	}
}

func ToKubernetesDeploymentPtr(x *entity.Deployment) *appsv1.Deployment {
	if x == nil {
		return nil
	}

	d := ToKubernetesDeployment(*x)

	return &d
}
