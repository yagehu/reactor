package reagent

import (
	"context"

	"github.com/yagehu/reactor/internal/errs"
	"github.com/yagehu/reactor/internal/model"
)

type DeleteReagentParams struct {
	Record model.Reagent
}

type DeleteReagentResult struct {
	Record model.Reagent
}

func (r *repository) DeleteReagent(
	ctx context.Context, p *DeleteReagentParams,
) (*DeleteReagentResult, error) {
	var op errs.Op = "repository/reagent.DeleteReagent"

	stmt, err := r.db.Prepare(`
UPDATE reagent
SET deleted_at = $1
WHERE guid = $2 AND deleted_at IS NULL
	`)
	if err != nil {
		return nil, errs.E(op, err)
	}

	now := r.nowFunc()
	record := p.Record
	record.DeletedAt = &now

	_, err = stmt.ExecContext(ctx, now, p.Record.GUID)
	if err != nil {
		return nil, errs.E(op, err)
	}

	return &DeleteReagentResult{
		Record: record,
	}, nil
}
