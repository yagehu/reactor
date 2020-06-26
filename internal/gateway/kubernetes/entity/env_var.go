package entity

type EnvVar struct {
	Name      string
	Value     string
	ValueFrom *EnvVarSource
}

func NewEnvVar(name, value string, valueFrom *EnvVarSource) EnvVar {
	return EnvVar{
		Name:      name,
		Value:     value,
		ValueFrom: valueFrom,
	}
}
