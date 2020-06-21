package repository

import (
	"go.uber.org/fx"

	reagentrepository "github.com/yagehu/reactor/internal/repository/reagent"
)

var Module = fx.Provide(reagentrepository.New)
