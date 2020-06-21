package nowfx

import (
	"time"

	"go.uber.org/fx"
)

var Module = fx.Provide(New)

func New() func() time.Time {
	return time.Now
}
