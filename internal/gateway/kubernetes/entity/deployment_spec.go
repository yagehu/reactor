package entity

type DeploymentSpec struct {
	Replicas *int32
	Selector *LabelSelector
	Template PodTemplateSpec
}

func NewDeploymentSpec(
	replicas *int32,
	selector *LabelSelector,
	template PodTemplateSpec,
) DeploymentSpec {
	return DeploymentSpec{
		Replicas: replicas,
		Selector: selector,
		Template: template,
	}
}
