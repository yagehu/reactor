package reagent

import (
	"context"

	"github.com/yagehu/reactor/internal/errs"
	"github.com/yagehu/reactor/internal/model"
)

type GetAllActiveReagentsParams struct {
}

type GetAllActiveReagentsResult struct {
	Records []model.Reagent
}

func (r *repository) GetAllActiveReagents(
	ctx context.Context, p *GetAllActiveReagentsParams,
) (*GetAllActiveReagentsResult, error) {
	var (
		op      errs.Op = "repository/reagent.GetAllActiveReagents"
		records []model.Reagent
	)

	rows, err := r.db.QueryContext(
		ctx,
		`
SELECT id, guid, created_at FROM reagent
WHERE deleted_at IS NULL
		`,
	)
	if err != nil {
		return nil, errs.E(op, err)
	}

	for rows.Next() {
		var record model.Reagent

		if err := rows.Scan(
			&record.ID,
			&record.GUID,
			&record.CreatedAt,
		); err != nil {
			return nil, errs.E(op, err)
		}

		records = append(records, record)
	}

	return &GetAllActiveReagentsResult{
		Records: records,
	}, nil
}
