package entity

type PodTemplateSpec struct {
	ObjectMeta

	Spec PodSpec
}

func NewPodTemplateSpec(objectMeta ObjectMeta, spec PodSpec) PodTemplateSpec {
	return PodTemplateSpec{
		ObjectMeta: objectMeta,
		Spec:       spec,
	}
}
