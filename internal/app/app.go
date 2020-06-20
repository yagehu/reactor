package app

import (
	"go.uber.org/fx"

	"github.com/yagehu/reactor/internal/fx/zapfx"
)

func New() fx.Option {
	return fx.Options(
		zapfx.Module,
	)
}
