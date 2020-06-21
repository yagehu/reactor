package kubernetes

import (
	"go.uber.org/fx"
)

type Gateway interface {
}

type Params struct {
	fx.In
}

func New(p Params) (Gateway, error) {
	return &gateway{}, nil
}

type gateway struct {
}
