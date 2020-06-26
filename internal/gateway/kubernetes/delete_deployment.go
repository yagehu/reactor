package kubernetes

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/yagehu/reactor/internal/errs"
)

type DeleteDeploymentParams struct {
	Namespace string
	Name      string
}

type DeleteDeploymentResult struct {
}

func (g *gateway) DeleteDeployment(
	ctx context.Context, p *DeleteDeploymentParams,
) (*DeleteDeploymentResult, error) {
	var op errs.Op = "gateway/kubernetes.DeleteDeployment"

	err := g.kubernetesClient.
		AppsV1().
		Deployments(p.Namespace).
		Delete(
			ctx,
			p.Name,
			metav1.DeleteOptions{},
		)
	if err != nil {
		return nil, errs.E(op, err)
	}

	return &DeleteDeploymentResult{}, nil
}
