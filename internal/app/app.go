package app

import (
	"go.uber.org/fx"

	"github.com/yagehu/reactor/config"
	"github.com/yagehu/reactor/internal/fx/zapfx"
)

func New() fx.Option {
	return fx.Options(
		config.Module,

		zapfx.Module,
	)
}
