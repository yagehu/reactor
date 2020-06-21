package reagent

import (
	"context"
	"database/sql"
	"time"

	"go.uber.org/fx"
)

type Repository interface {
	CreateReagent(
		context.Context, *CreateReagentParams,
	) (*CreateReagentResult, error)
	GetAllReagents(
		context.Context, *GetAllReagentsParams,
	) (*GetAllReagentsResult, error)
}

type Params struct {
	fx.In

	Db      *sql.DB
	NowFunc func() time.Time
}

func New(p Params) (Repository, error) {
	return &repository{
		db:      p.Db,
		nowFunc: p.NowFunc,
	}, nil
}

type repository struct {
	db      *sql.DB
	nowFunc func() time.Time
}
