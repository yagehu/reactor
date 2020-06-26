package mapper

import (
	v1 "k8s.io/api/core/v1"

	"github.com/yagehu/reactor/internal/gateway/kubernetes/entity"
)

func ToKubernetesPodTemplateSpec(x entity.PodTemplateSpec) v1.PodTemplateSpec {
	return v1.PodTemplateSpec{
		ObjectMeta: ToKubernetesObjectMeta(x.ObjectMeta),
		Spec:       ToKubernetesPodSpec(x.Spec),
	}
}
