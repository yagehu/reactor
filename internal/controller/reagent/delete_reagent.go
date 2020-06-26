package reagent

import (
	"context"

	"github.com/yagehu/reactor/internal/entity"
	"github.com/yagehu/reactor/internal/errs"
	kubernetesgateway "github.com/yagehu/reactor/internal/gateway/kubernetes"
	"github.com/yagehu/reactor/internal/mapper"
	reagentrepository "github.com/yagehu/reactor/internal/repository/reagent"
)

type DeleteReagentParams struct {
	Reagent entity.Reagent
}

type DeleteReagentResult struct {
}

func (c *controller) DeleteReagent(
	ctx context.Context, p *DeleteReagentParams,
) (*DeleteReagentResult, error) {
	var op errs.Op = "controller/reagent.DeleteReagent"

	_, err := c.kubernetesGateway.DeleteDeployment(
		ctx,
		&kubernetesgateway.DeleteDeploymentParams{
			Namespace: c.config.Reactor.Namespace,
			Name:      c.config.Reactor.Product.Name + "." + p.Reagent.ID.String(),
		},
	)
	if err != nil {
		return nil, errs.E(op, err)
	}

	_, err = c.reagentRepository.DeleteReagent(
		ctx,
		&reagentrepository.DeleteReagentParams{
			Record: mapper.ToReagentModel(p.Reagent),
		},
	)
	if err != nil {
		return nil, errs.E(op, err)
	}

	return &DeleteReagentResult{}, nil
}
