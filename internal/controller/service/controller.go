package service

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/yagehu/reactor/config"
	reagentcontroller "github.com/yagehu/reactor/internal/controller/reagent"
)

type Controller interface {
	ProcessService(
		context.Context, *ProcessServiceParams,
	) (*ProcessServiceResult, error)
	ProcessServices(
		context.Context, *ProcessServicesParams,
	) (*ProcessServicesResult, error)
}

type Params struct {
	fx.In

	Config            config.Config
	Logger            *zap.Logger
	ReagentController reagentcontroller.Controller
}

func New(p Params) (Controller, error) {
	return &controller{
		config:            p.Config,
		logger:            p.Logger,
		reagentController: p.ReagentController,
	}, nil
}

type controller struct {
	config            config.Config
	logger            *zap.Logger
	reagentController reagentcontroller.Controller
}
