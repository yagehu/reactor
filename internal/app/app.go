package app

import (
	"go.uber.org/fx"

	"github.com/yagehu/reactor/config"
	"github.com/yagehu/reactor/internal/controller"
	"github.com/yagehu/reactor/internal/fx/dbfx"
	"github.com/yagehu/reactor/internal/fx/httpfx"
	"github.com/yagehu/reactor/internal/fx/middlewarefx"
	"github.com/yagehu/reactor/internal/fx/uuidfx"
	"github.com/yagehu/reactor/internal/fx/zapfx"
	"github.com/yagehu/reactor/internal/handler"
	"github.com/yagehu/reactor/internal/repository"
)

func New() fx.Option {
	return fx.Options(
		config.Module,
		controller.Module,
		handler.Module,
		repository.Module,

		dbfx.Module,
		httpfx.Module,
		middlewarefx.Module,
		uuidfx.Module,
		zapfx.Module,
	)
}
