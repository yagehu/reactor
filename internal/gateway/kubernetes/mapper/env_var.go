package mapper

import (
	v1 "k8s.io/api/core/v1"

	"github.com/yagehu/reactor/internal/gateway/kubernetes/entity"
)

func ToKubernetesEnvVar(x entity.EnvVar) v1.EnvVar {
	return v1.EnvVar{
		Name:      x.Name,
		Value:     x.Value,
		ValueFrom: ToKubernetesEnvVarSourcePtr(x.ValueFrom),
	}
}

func ToKubernetesEnvVarList(xs []entity.EnvVar) []v1.EnvVar {
	s := make([]v1.EnvVar, 0, len(xs))

	for _, x := range xs {
		s = append(s, ToKubernetesEnvVar(x))
	}

	return s
}
