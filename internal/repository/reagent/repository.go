package reagent

import (
	"context"
	"database/sql"

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

	Db *sql.DB
}

func New(p Params) (Repository, error) {
	return &repository{
		db: p.Db,
	}, nil
}

type repository struct {
	db *sql.DB
}
