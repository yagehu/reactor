package app

import (
	"go.uber.org/fx"

	"github.com/yagehu/reactor/config"
	"github.com/yagehu/reactor/internal/controller"
	"github.com/yagehu/reactor/internal/fx/dbfx"
	"github.com/yagehu/reactor/internal/fx/httpfx"
	"github.com/yagehu/reactor/internal/fx/kubernetesfx"
	"github.com/yagehu/reactor/internal/fx/middlewarefx"
	"github.com/yagehu/reactor/internal/fx/nowfx"
	"github.com/yagehu/reactor/internal/fx/uuidfx"
	"github.com/yagehu/reactor/internal/fx/zapfx"
	"github.com/yagehu/reactor/internal/gateway"
	"github.com/yagehu/reactor/internal/handler"
	"github.com/yagehu/reactor/internal/repository"
)

func New() fx.Option {
	return fx.Options(
		config.Module,
		controller.Module,
		gateway.Module,
		handler.Module,
		repository.Module,

		dbfx.Module,
		httpfx.Module,
		kubernetesfx.Module,
		middlewarefx.Module,
		nowfx.Module,
		uuidfx.Module,
		zapfx.Module,
	)
}
