package kubernetes

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/yagehu/reactor/internal/errs"
	"github.com/yagehu/reactor/internal/gateway/kubernetes/entity"
	"github.com/yagehu/reactor/internal/gateway/kubernetes/mapper"
)

type CreateDeploymentParams struct {
	Namespace  string
	Deployment *entity.Deployment
}

type CreateDeploymentResult struct {
}

func (g *gateway) CreateDeployment(
	ctx context.Context, p *CreateDeploymentParams,
) (*CreateDeploymentResult, error) {
	var op errs.Op = "gateway/kubernetes.CreateDeployment"

	_, err := g.kubernetesClient.
		AppsV1().
		Deployments(p.Namespace).
		Create(
			ctx,
			mapper.ToKubernetesDeploymentPtr(p.Deployment),
			metav1.CreateOptions{},
		)
	if err != nil {
		return nil, errs.E(op, err)
	}

	return &CreateDeploymentResult{}, nil
}
