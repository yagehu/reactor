package reactor

import (
	"context"

	"go.uber.org/fx"

	reagentcontroller "github.com/yagehu/reactor/internal/controller/reagent"
)

type Controller interface {
	WatchEvent(context.Context, *WatchEventParams) (*WatchEventResult, error)
}

type Params struct {
	fx.In

	ReagentController reagentcontroller.Controller
}

func New(p Params) (Controller, error) {
	return &controller{
		reagentController: p.ReagentController,
	}, nil
}

type controller struct {
	reagentController reagentcontroller.Controller
}
