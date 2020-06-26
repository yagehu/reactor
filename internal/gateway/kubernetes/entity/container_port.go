package entity

type ContainerPort struct {
	Name          string
	HostPort      int32
	ContainerPort int32
	Protocol      Protocol
	HostIP        string
}

func NewContainerPort(
	name string,
	hostPort int32,
	containerPort int32,
	protocol Protocol,
	hostIP string,
) ContainerPort {
	return ContainerPort{
		Name:          name,
		HostPort:      hostPort,
		ContainerPort: containerPort,
		Protocol:      protocol,
		HostIP:        hostIP,
	}
}
