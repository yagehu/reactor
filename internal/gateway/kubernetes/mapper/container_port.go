package mapper

import (
	v1 "k8s.io/api/core/v1"

	"github.com/yagehu/reactor/internal/gateway/kubernetes/entity"
)

func ToKubernetesContainerPort(x entity.ContainerPort) v1.ContainerPort {
	return v1.ContainerPort{
		Name:          x.Name,
		HostPort:      x.HostPort,
		ContainerPort: x.ContainerPort,
		Protocol:      v1.Protocol(x.Protocol.String()),
		HostIP:        x.HostIP,
	}
}

func ToKubernetesContainerPortList(
	xs []entity.ContainerPort,
) []v1.ContainerPort {
	s := make([]v1.ContainerPort, 0, len(xs))

	for _, x := range xs {
		s = append(s, ToKubernetesContainerPort(x))
	}

	return s
}
