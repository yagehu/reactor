package reactor

import (
	"context"

	"github.com/gofrs/uuid"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/yagehu/reactor/config"
	reagentcontroller "github.com/yagehu/reactor/internal/controller/reagent"
)

type Controller interface {
	WatchEvent(context.Context, *WatchEventParams) (*WatchEventResult, error)
}

type Params struct {
	fx.In

	Config            config.Config
	Logger            *zap.Logger
	UUIDGenerator     uuid.Generator
	ReagentController reagentcontroller.Controller
}

func New(p Params) (Controller, error) {
	return &controller{
		config:            p.Config,
		logger:            p.Logger,
		uuidGenerator:     p.UUIDGenerator,
		reagentController: p.ReagentController,
	}, nil
}

type controller struct {
	config            config.Config
	logger            *zap.Logger
	uuidGenerator     uuid.Generator
	reagentController reagentcontroller.Controller
}
