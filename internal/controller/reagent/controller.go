package reagent

import (
	"context"

	"github.com/gofrs/uuid"
	"go.uber.org/fx"

	reagentrepository "github.com/yagehu/reactor/internal/repository/reagent"
)

type Controller interface {
	CreateReagent(
		context.Context, *CreateReagentParams,
	) (*CreateReagentResult, error)
}

type Params struct {
	fx.In

	UUIDGenerator     uuid.Generator
	ReagentRepository reagentrepository.Repository
}

func New(p Params) (Controller, error) {
	return &controller{
		uuidGenerator:     p.UUIDGenerator,
		reagentRepository: p.ReagentRepository,
	}, nil
}

type controller struct {
	uuidGenerator     uuid.Generator
	reagentRepository reagentrepository.Repository
}
