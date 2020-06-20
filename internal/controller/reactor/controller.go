package reactor

import (
	"go.uber.org/fx"
)

type Controller interface {
}

type Params struct {
	fx.In
}

func New(p Params) (Controller, error) {
	return &controller{}, nil
}

type controller struct {
}
