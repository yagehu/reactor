package reagent

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"go.uber.org/fx"

	reagentrepository "github.com/yagehu/reactor/internal/repository/reagent"
)

type Controller interface {
	CreateReagent(
		context.Context, *CreateReagentParams,
	) (*CreateReagentResult, error)
	GetAllReagents(
		context.Context, *GetAllReagentsParams,
	) (*GetAllReagentsResult, error)
}

type Params struct {
	fx.In

	NowFunc           func() time.Time
	UUIDGenerator     uuid.Generator
	ReagentRepository reagentrepository.Repository
}

func New(p Params) (Controller, error) {
	return &controller{
		nowFunc:           p.NowFunc,
		uuidGenerator:     p.UUIDGenerator,
		reagentRepository: p.ReagentRepository,
	}, nil
}

type controller struct {
	nowFunc           func() time.Time
	uuidGenerator     uuid.Generator
	reagentRepository reagentrepository.Repository
}
