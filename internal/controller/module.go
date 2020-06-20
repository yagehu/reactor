package controller

import (
	"go.uber.org/fx"

	"github.com/yagehu/reactor/internal/controller/reactor"
)

var Module = fx.Provide(reactor.New)
