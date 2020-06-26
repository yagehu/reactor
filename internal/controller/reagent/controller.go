package reagent

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/yagehu/reactor/config"
	kubernetesgateway "github.com/yagehu/reactor/internal/gateway/kubernetes"
	reagentrepository "github.com/yagehu/reactor/internal/repository/reagent"
)

type Controller interface {
	CreateReagent(
		context.Context, *CreateReagentParams,
	) (*CreateReagentResult, error)
	DeleteReagent(
		context.Context, *DeleteReagentParams,
	) (*DeleteReagentResult, error)
	GetAllReagents(
		context.Context, *GetAllReagentsParams,
	) (*GetAllReagentsResult, error)
	GetReagentByGUID(
		context.Context, *GetReagentByGUIDParams,
	) (*GetReagentByGUIDResult, error)
	Reconcile(context.Context, *ReconcileParams) (*ReconcileResult, error)
}

type Params struct {
	fx.In

	Config            config.Config
	Logger            *zap.Logger
	NowFunc           func() time.Time
	UUIDGenerator     uuid.Generator
	KubernetesGateway kubernetesgateway.Gateway
	ReagentRepository reagentrepository.Repository
}

func New(p Params) (Controller, error) {
	return &controller{
		config:            p.Config,
		logger:            p.Logger,
		nowFunc:           p.NowFunc,
		uuidGenerator:     p.UUIDGenerator,
		kubernetesGateway: p.KubernetesGateway,
		reagentRepository: p.ReagentRepository,
	}, nil
}

type controller struct {
	config            config.Config
	logger            *zap.Logger
	nowFunc           func() time.Time
	uuidGenerator     uuid.Generator
	kubernetesGateway kubernetesgateway.Gateway
	reagentRepository reagentrepository.Repository
}
