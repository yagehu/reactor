package entity

type ObjectMeta struct {
	Name      string
	Namespace string
	Labels    map[string]string
}

func NewObjectMeta(
	name, namespace string, labels map[string]string,
) ObjectMeta {
	return ObjectMeta{
		Name:      name,
		Namespace: namespace,
		Labels:    labels,
	}
}
