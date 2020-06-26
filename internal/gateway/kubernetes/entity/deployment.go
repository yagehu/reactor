package entity

type Deployment struct {
	TypeMeta
	ObjectMeta

	Spec DeploymentSpec
}

func NewDeployment(
	typeMeta TypeMeta, objectMeta ObjectMeta, spec DeploymentSpec,
) Deployment {
	return Deployment{
		TypeMeta:   typeMeta,
		ObjectMeta: objectMeta,
		Spec:       spec,
	}
}

func DeploymentPtr(d Deployment) *Deployment {
	return &d
}
