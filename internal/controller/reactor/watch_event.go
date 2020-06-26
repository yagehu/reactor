package reactor

import (
	"context"

	"golang.org/x/sync/errgroup"

	servicecontroller "github.com/yagehu/reactor/internal/controller/service"
	"github.com/yagehu/reactor/internal/entity"
	"github.com/yagehu/reactor/internal/errs"
)

type WatchEventParams struct {
	Services []entity.Service
}

type WatchEventResult struct {
}

func (c *controller) WatchEvent(
	ctx context.Context, p *WatchEventParams,
) (*WatchEventResult, error) {
	var (
		op       errs.Op = "controller/reactor.WatchEvent"
		reagents []entity.Reagent
	)

	eg, ectx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		res, err := c.serviceController.ProcessServices(
			ectx,
			&servicecontroller.ProcessServicesParams{
				Services: p.Services,
			},
		)
		if err != nil {
			return err
		}

		reagents = res.Reagents

		return nil
	})

	if err := eg.Wait(); err != nil {
		return nil, errs.E(op, err)
	}

	return &WatchEventResult{}, nil
}
