package entity

type ObjectFieldSelector struct {
	APIVersion string
	FieldPath  string
}

func NewObjectFieldSelector(
	apiVersion string, fieldPath string,
) ObjectFieldSelector {
	return ObjectFieldSelector{
		APIVersion: apiVersion,
		FieldPath:  fieldPath,
	}
}
