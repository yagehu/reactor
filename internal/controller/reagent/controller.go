package reagent

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"go.uber.org/fx"

	"github.com/yagehu/reactor/config"
	kubernetesgateway "github.com/yagehu/reactor/internal/gateway/kubernetes"
	reagentrepository "github.com/yagehu/reactor/internal/repository/reagent"
)

type Controller interface {
	CreateReagent(
		context.Context, *CreateReagentParams,
	) (*CreateReagentResult, error)
	GetAllReagents(
		context.Context, *GetAllReagentsParams,
	) (*GetAllReagentsResult, error)
	GetReagentByGUID(
		context.Context, *GetReagentByGUIDParams,
	) (*GetReagentByGUIDResult, error)
}

type Params struct {
	fx.In

	Config            config.Config
	NowFunc           func() time.Time
	UUIDGenerator     uuid.Generator
	KubernetesGateway kubernetesgateway.Gateway
	ReagentRepository reagentrepository.Repository
}

func New(p Params) (Controller, error) {
	return &controller{
		config:            p.Config,
		nowFunc:           p.NowFunc,
		uuidGenerator:     p.UUIDGenerator,
		kubernetesGateway: p.KubernetesGateway,
		reagentRepository: p.ReagentRepository,
	}, nil
}

type controller struct {
	config            config.Config
	nowFunc           func() time.Time
	uuidGenerator     uuid.Generator
	kubernetesGateway kubernetesgateway.Gateway
	reagentRepository reagentrepository.Repository
}
