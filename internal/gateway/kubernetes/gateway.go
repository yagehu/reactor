package kubernetes

import (
	"context"

	"go.uber.org/fx"
	"k8s.io/client-go/kubernetes"
)

type Gateway interface {
	CreateDeployment(
		context.Context, *CreateDeploymentParams,
	) (*CreateDeploymentResult, error)
	DeleteDeployment(
		context.Context, *DeleteDeploymentParams,
	) (*DeleteDeploymentResult, error)
}

type Params struct {
	fx.In

	KubernetesClient kubernetes.Interface
}

func New(p Params) (Gateway, error) {
	return &gateway{
		kubernetesClient: p.KubernetesClient,
	}, nil
}

type gateway struct {
	kubernetesClient kubernetes.Interface
}
