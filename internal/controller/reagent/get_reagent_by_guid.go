package reagent

import (
	"context"

	"github.com/gofrs/uuid"

	"github.com/yagehu/reactor/internal/entity"
	"github.com/yagehu/reactor/internal/errs"
	"github.com/yagehu/reactor/internal/mapper"
	reagentrepository "github.com/yagehu/reactor/internal/repository/reagent"
)

type GetReagentByGUIDParams struct {
	GUID uuid.UUID
}

type GetReagentByGUIDResult struct {
	Reagent entity.Reagent
}

func (c *controller) GetReagentByGUID(
	ctx context.Context, p *GetReagentByGUIDParams,
) (*GetReagentByGUIDResult, error) {
	var op errs.Op = "controller/reagent.GetReagentByGUID"

	res, err := c.reagentRepository.GetReagentByGUID(
		ctx,
		&reagentrepository.GetReagentByGUIDParams{GUID: p.GUID.String()},
	)
	if err != nil {
		return nil, errs.E(op, err)
	}

	reagent, err := mapper.FromReagentModel(res.Record)
	if err != nil {
		return nil, errs.E(op, err)
	}

	return &GetReagentByGUIDResult{
		Reagent: reagent,
	}, nil
}
