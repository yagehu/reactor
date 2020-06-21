package reagent

import (
	"context"
	"database/sql"
	"time"

	"github.com/yagehu/reactor/internal/errs"
	"github.com/yagehu/reactor/internal/model"
)

type GetReagentByGUIDParams struct {
	GUID string
}

type GetReagentByGUIDResult struct {
	Record model.Reagent
}

func (r *repository) GetReagentByGUID(
	ctx context.Context, p *GetReagentByGUIDParams,
) (*GetReagentByGUIDResult, error) {
	var (
		op        errs.Op = "repository/reagent.GetReagentByGUID"
		id        string
		guid      string
		createdAt time.Time
		deletedAt *time.Time
	)

	err := r.db.QueryRowContext(
		ctx,
		`
SELECT id, guid, created_at, deleted_at
FROM reagent
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1
		`,
		p.GUID,
	).Scan(&id, &guid, &createdAt, &deletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.E(op, errs.KindReagentNotFound, err)
		}

		return nil, errs.E(op, err)
	}

	return &GetReagentByGUIDResult{
		Record: model.NewReagent(id, guid, createdAt, deletedAt),
	}, nil
}
