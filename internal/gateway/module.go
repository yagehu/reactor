package gateway

import (
	"go.uber.org/fx"

	kubernetesgateway "github.com/yagehu/reactor/internal/gateway/kubernetes"
)

var Module = fx.Provide(
	kubernetesgateway.New,
)
