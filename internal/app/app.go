package app

import (
	"go.uber.org/fx"

	"github.com/yagehu/reactor/config"
	"github.com/yagehu/reactor/internal/controller"
	"github.com/yagehu/reactor/internal/fx/httpfx"
	"github.com/yagehu/reactor/internal/fx/middlewarefx"
	"github.com/yagehu/reactor/internal/fx/zapfx"
	"github.com/yagehu/reactor/internal/handler"
)

func New() fx.Option {
	return fx.Options(
		config.Module,
		controller.Module,
		handler.Module,

		httpfx.Module,
		middlewarefx.Module,
		zapfx.Module,
	)
}
