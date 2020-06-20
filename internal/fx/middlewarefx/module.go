package middlewarefx

import (
	"go.uber.org/fx"

	"github.com/yagehu/reactor/internal/fx/middlewarefx/logging"
)

var Module = fx.Provide(
	logging.New,
)
