package mapper

import (
	appsv1 "k8s.io/api/apps/v1"

	"github.com/yagehu/reactor/internal/gateway/kubernetes/entity"
)

func ToKubernetesDeploymentSpec(x entity.DeploymentSpec) appsv1.DeploymentSpec {
	return appsv1.DeploymentSpec{
		Replicas: x.Replicas,
		Selector: ToKubernetesLabelSelectorPtr(x.Selector),
		Template: ToKubernetesPodTemplateSpec(x.Template),
	}
}
