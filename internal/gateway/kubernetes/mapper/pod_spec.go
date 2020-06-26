package mapper

import (
	v1 "k8s.io/api/core/v1"

	"github.com/yagehu/reactor/internal/gateway/kubernetes/entity"
)

func ToKubernetesPodSpec(x entity.PodSpec) v1.PodSpec {
	return v1.PodSpec{
		Containers:    ToKubernetesContainerList(x.Containers),
		RestartPolicy: v1.RestartPolicy(x.RestartPolicy),
	}
}
