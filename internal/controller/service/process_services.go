package service

import (
	"context"
	"sync"

	"go.uber.org/zap"

	"github.com/yagehu/reactor/internal/entity"
	"github.com/yagehu/reactor/internal/errs"
)

type ProcessServicesParams struct {
	Services []entity.Service
}

type ProcessServicesResult struct {
	Reagents []entity.Reagent
}

func (c *controller) ProcessServices(
	ctx context.Context, p *ProcessServicesParams,
) (*ProcessServicesResult, error) {
	var (
		wg       sync.WaitGroup
		mu       sync.Mutex
		reagents []entity.Reagent
	)

	for _, service := range p.Services {
		wg.Add(1)

		go func(service entity.Service) {
			defer wg.Done()

			res, err := c.ProcessService(
				ctx, &ProcessServiceParams{Service: service},
			)
			if err != nil {
				e, ok := err.(*errs.Error)
				if ok && e.Kind == errs.KindInvalidService {
					// Do not log if it's an invalid service.
					return
				}

				c.logger.Error("Could not process service.", zap.Error(err))
				return
			}

			mu.Lock()
			reagents = append(reagents, res.Reagent)
			mu.Unlock()
		}(service)
	}

	wg.Wait()

	return &ProcessServicesResult{}, nil
}
