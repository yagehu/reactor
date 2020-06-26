package service

import (
	"context"

	"github.com/gofrs/uuid"

	reagentcontroller "github.com/yagehu/reactor/internal/controller/reagent"
	"github.com/yagehu/reactor/internal/entity"
	"github.com/yagehu/reactor/internal/errs"
)

type ProcessServiceParams struct {
	Service entity.Service
}

type ProcessServiceResult struct {
	Reagent entity.Reagent
}

func (c *controller) ProcessService(
	ctx context.Context, p *ProcessServiceParams,
) (*ProcessServiceResult, error) {
	var op errs.Op = "controller/service.ProcessService"

	if _, ok := p.Service.Tags[c.config.Reactor.ReactTo]; !ok {
		return nil, errs.E(op, errs.KindInvalidService)
	}

	guid, err := uuid.FromString(p.Service.Name)
	if err != nil {
		return nil, errs.E(op, err)
	}

	var reagent entity.Reagent

	res, err := c.reagentController.GetReagentByGUID(
		ctx,
		&reagentcontroller.GetReagentByGUIDParams{GUID: guid},
	)
	if err != nil {
		e, ok := err.(*errs.Error)
		if !ok || e.Kind != errs.KindReagentNotFound {
			return nil, errs.E(op, err)
		}

		res, err := c.reagentController.CreateReagent(
			ctx,
			&reagentcontroller.CreateReagentParams{GUID: guid},
		)
		if err != nil {
			return nil, errs.E(op, err)
		}

		reagent = res.Reagent
	} else {
		reagent = res.Reagent
	}

	return &ProcessServiceResult{Reagent: reagent}, nil
}
