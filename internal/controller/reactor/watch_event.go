package reactor

import (
	"context"

	"github.com/yagehu/reactor/internal/entity"
)

type WatchEventParams struct {
	Services []entity.Service
}

type WatchEventResult struct {
}

func (c *controller) WatchEvent(
	ctx context.Context, p *WatchEventParams,
) (*WatchEventResult, error) {
	return &WatchEventResult{}, nil
}
