package zapfx

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In
}

type Result struct {
	fx.Out

	Logger *zap.Logger
}

func New(p Params) (Result, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return Result{}, err
	}

	return Result{
		Logger: logger,
	}, nil
}
