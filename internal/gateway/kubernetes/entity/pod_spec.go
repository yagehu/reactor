package entity

type PodSpec struct {
	Containers    []Container
	RestartPolicy RestartPolicy
}

func NewPodSpec(containers []Container, restartPolicy RestartPolicy) PodSpec {
	return PodSpec{
		Containers:    containers,
		RestartPolicy: restartPolicy,
	}
}
