package mapper

import (
	v1 "k8s.io/api/core/v1"

	"github.com/yagehu/reactor/internal/gateway/kubernetes/entity"
)

func ToKubernetesEnvVarSource(x entity.EnvVarSource) v1.EnvVarSource {
	return v1.EnvVarSource{
		FieldRef: ToKubernetesObjectFieldSelectorPtr(x.FieldRef),
	}
}

func ToKubernetesEnvVarSourcePtr(x *entity.EnvVarSource) *v1.EnvVarSource {
	if x == nil {
		return nil
	}

	envVarSource := ToKubernetesEnvVarSource(*x)

	return &envVarSource
}
