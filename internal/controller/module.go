package controller

import (
	"go.uber.org/fx"

	reactorcontroller "github.com/yagehu/reactor/internal/controller/reactor"
	reagentcontroller "github.com/yagehu/reactor/internal/controller/reagent"
)

var Module = fx.Provide(
	reactorcontroller.New,
	reagentcontroller.New,
)
