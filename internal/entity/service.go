package entity

type Service struct {
	Name string
	Tags map[string]struct{}
}

func NewService(name string, tags map[string]struct{}) Service {
	return Service{
		Name: name,
		Tags: tags,
	}
}
