package reagent

import (
	"context"

	"github.com/gofrs/uuid"

	"github.com/yagehu/reactor/internal/entity"
	"github.com/yagehu/reactor/internal/errs"
	kubernetesgateway "github.com/yagehu/reactor/internal/gateway/kubernetes"
	kubernetesentity "github.com/yagehu/reactor/internal/gateway/kubernetes/entity"
	"github.com/yagehu/reactor/internal/mapper"
	"github.com/yagehu/reactor/internal/ptr"
	reagentrepository "github.com/yagehu/reactor/internal/repository/reagent"
)

type CreateReagentParams struct {
	GUID uuid.UUID
}

type CreateReagentResult struct {
	Reagent entity.Reagent
}

func (c *controller) CreateReagent(
	ctx context.Context, p *CreateReagentParams,
) (*CreateReagentResult, error) {
	var op errs.Op = "controller/reagent.CreateReagent"

	id, err := c.uuidGenerator.NewV4()
	if err != nil {
		return nil, errs.E(op, err)
	}

	_, err = c.kubernetesGateway.CreateDeployment(
		ctx,
		&kubernetesgateway.CreateDeploymentParams{
			Namespace: c.config.Reactor.Namespace,
			Deployment: kubernetesentity.DeploymentPtr(
				kubernetesentity.NewDeployment(
					kubernetesentity.NewTypeMeta(
						"Deployment",
						"apps/v1",
					),
					kubernetesentity.NewObjectMeta(
						c.config.Reactor.Product.Name+"."+id.String(),
						c.config.Reactor.Namespace,
						nil,
					),
					kubernetesentity.NewDeploymentSpec(
						ptr.Int32Ptr(1),
						kubernetesentity.LabelSelectorPtr(
							kubernetesentity.NewLabelSelector(
								map[string]string{
									"app": c.config.Reactor.Product.Name,
									"id":  id.String(),
								},
								nil,
							),
						),
						kubernetesentity.NewPodTemplateSpec(
							kubernetesentity.NewObjectMeta(
								c.config.Reactor.Product.Name,
								c.config.Reactor.Namespace,
								map[string]string{
									"app": c.config.Reactor.Product.Name,
									"id":  id.String(),
								},
							),
							kubernetesentity.NewPodSpec(
								[]kubernetesentity.Container{
									kubernetesentity.NewContainer(
										c.config.Reactor.Product.Name,
										c.config.Reactor.Product.Image,
										nil,
										nil,
										nil,
										nil,
									),
								},
								kubernetesentity.RestartPolicyAlways,
							),
						),
					),
				),
			),
		},
	)
	if err != nil {
		return nil, errs.E(op, err)
	}

	reagent := entity.NewReagent(
		entity.NewReagentID(id), p.GUID, c.nowFunc(), nil,
	)

	_, err = c.reagentRepository.CreateReagent(
		ctx,
		&reagentrepository.CreateReagentParams{
			Record: mapper.ToReagentModel(reagent),
		},
	)
	if err != nil {
		return nil, errs.E(op, err)
	}

	return &CreateReagentResult{Reagent: reagent}, nil
}
