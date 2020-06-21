package reagent

import (
	"context"

	"github.com/yagehu/reactor/internal/entity"
	"github.com/yagehu/reactor/internal/errs"
	"github.com/yagehu/reactor/internal/mapper"
	reagentrepository "github.com/yagehu/reactor/internal/repository/reagent"
)

type GetAllReagentsParams struct {
}

type GetAllReagentsResult struct {
	Reagents []entity.Reagent
}

func (c *controller) GetAllReagents(
	ctx context.Context, _ *GetAllReagentsParams,
) (*GetAllReagentsResult, error) {
	var op errs.Op = "controller/reagent.GetAllReagents"

	res, err := c.reagentRepository.GetAllReagents(
		ctx, &reagentrepository.GetAllReagentsParams{},
	)
	if err != nil {
		return nil, errs.E(op, err)
	}

	reagents, err := mapper.FromReagentModelList(res.Records)
	if err != nil {
		return nil, errs.E(op, err)
	}

	return &GetAllReagentsResult{Reagents: reagents}, nil
}
