package entity

type TypeMeta struct {
	Kind       string
	APIVersion string
}

func NewTypeMeta(kind string, apiVersion string) TypeMeta {
	return TypeMeta{
		Kind:       kind,
		APIVersion: apiVersion,
	}
}
