package reactor

import (
	"context"

	"go.uber.org/fx"
)

type Controller interface {
	WatchEvent(context.Context, *WatchEventParams) (*WatchEventResult, error)
}

type Params struct {
	fx.In
}

func New(p Params) (Controller, error) {
	return &controller{}, nil
}

type controller struct {
}
