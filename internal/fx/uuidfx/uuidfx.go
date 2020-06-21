package uuidfx

import (
	"github.com/gofrs/uuid"
	"go.uber.org/fx"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In
}

type Result struct {
	fx.Out

	Generator uuid.Generator
}

func New(p Params) (Result, error) {
	return Result{
		Generator: uuid.NewGen(),
	}, nil
}
