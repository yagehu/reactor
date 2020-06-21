package reactor

import (
	"context"

	"github.com/gofrs/uuid"

	reagentcontroller "github.com/yagehu/reactor/internal/controller/reagent"
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
	var reagents []entity.Reagent

	for _, service := range p.Services {
		if _, ok := service.Tags[c.config.Reactor.ReactTo]; !ok {
			continue
		}

		guid, err := uuid.FromString(service.Name)
		if err != nil {
			continue
		}

		var reagent entity.Reagent

		res, err := c.reagentController.GetReagentByGUID(
			ctx,
			&reagentcontroller.GetReagentByGUIDParams{GUID: guid},
		)
		if err != nil {
			e, ok := err.(*errs.Error)
			if !ok || e.Kind != errs.KindReagentNotFound {
				continue
			}

			res, err := c.reagentController.CreateReagent(
				ctx,
				&reagentcontroller.CreateReagentParams{GUID: guid},
			)
			if err != nil {
				continue
			}

			reagent = res.Reagent
		} else {
			reagent = res.Reagent
		}

		reagents = append(reagents, reagent)
	}

	return &WatchEventResult{}, nil
}
