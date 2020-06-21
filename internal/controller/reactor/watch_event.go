package reactor

import "context"

type WatchEventParams struct {
}

type WatchEventResult struct {
}

func (c *controller) WatchEvent(
	ctx context.Context, p *WatchEventParams,
) (*WatchEventResult, error) {
	return &WatchEventResult{}, nil
}
