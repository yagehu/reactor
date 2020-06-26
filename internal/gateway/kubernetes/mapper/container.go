package mapper

import (
	v1 "k8s.io/api/core/v1"

	"github.com/yagehu/reactor/internal/gateway/kubernetes/entity"
)

func ToKubernetesContainer(x entity.Container) v1.Container {
	return v1.Container{
		Name:    x.Name,
		Image:   x.Image,
		Command: x.Command,
		Args:    x.Args,
		Ports:   ToKubernetesContainerPortList(x.Ports),
		Env:     ToKubernetesEnvVarList(x.Env),
	}
}

func ToKubernetesContainerList(xs []entity.Container) []v1.Container {
	s := make([]v1.Container, 0, len(xs))

	for _, x := range xs {
		s = append(s, ToKubernetesContainer(x))
	}

	return s
}
