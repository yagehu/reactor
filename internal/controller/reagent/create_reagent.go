package reagent

import (
	"context"

	"github.com/gofrs/uuid"

	"github.com/yagehu/reactor/internal/entity"
	"github.com/yagehu/reactor/internal/errs"
	"github.com/yagehu/reactor/internal/mapper"
	reagentrepository "github.com/yagehu/reactor/internal/repository/reagent"
)

type CreateReagentParams struct {
	GUID uuid.UUID
}

type CreateReagentResult struct {
	Reagent entity.Reagent
}

func (c *controller) CreateReagent(
	ctx context.Context, p *CreateReagentParams,
) (*CreateReagentResult, error) {
	var op errs.Op = "controller/reagent.CreateReagent"

	id, err := c.uuidGenerator.NewV4()
	if err != nil {
		return nil, errs.E(op, err)
	}

	reagent := entity.NewReagent(id, p.GUID, c.nowFunc(), nil)

	_, err = c.reagentRepository.CreateReagent(
		ctx,
		&reagentrepository.CreateReagentParams{
			Record: mapper.ToReagentModel(reagent),
		},
	)
	if err != nil {
		return nil, errs.E(op, err)
	}

	return &CreateReagentResult{Reagent: reagent}, nil
}
