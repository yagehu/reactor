package reagent

import (
	"context"

	"github.com/yagehu/reactor/internal/errs"
	"github.com/yagehu/reactor/internal/model"
)

type CreateReagentParams struct {
	Record model.Reagent
}

type CreateReagentResult struct {
}

func (r *repository) CreateReagent(
	ctx context.Context, p *CreateReagentParams,
) (*CreateReagentResult, error) {
	var op errs.Op = "repository/reagent.CreateReagent"

	stmt, err := r.db.Prepare(`
INSERT INTO reagent (id, guid, created_at)
VALUES ($1, $2, $3)
	`)
	if err != nil {
		return nil, errs.E(op, err)
	}

	_, err = stmt.ExecContext(
		ctx, p.Record.ID, p.Record.GUID, p.Record.CreatedAt,
	)
	if err != nil {
		return nil, errs.E(op, err)
	}

	return &CreateReagentResult{}, nil
}
