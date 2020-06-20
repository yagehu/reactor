package handler

import (
	"go.uber.org/fx"

	"github.com/yagehu/reactor/internal/handler/reactor"
)

var Module = fx.Invoke(reactor.Register)
