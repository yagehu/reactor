package entity

type EnvVarSource struct {
	FieldRef *ObjectFieldSelector
}

func NewEnvVarSource(fieldRef *ObjectFieldSelector) EnvVarSource {
	return EnvVarSource{FieldRef: fieldRef}
}
