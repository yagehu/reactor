package entity

type Container struct {
	Name    string
	Image   string
	Command []string
	Args    []string
	Ports   []ContainerPort
	Env     []EnvVar
}

func NewContainer(
	name,
	image string,
	command,
	args []string,
	ports []ContainerPort,
	env []EnvVar,
) Container {
	return Container{
		Name:    name,
		Image:   image,
		Command: command,
		Args:    args,
		Ports:   ports,
		Env:     env,
	}
}
