package mapper

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/yagehu/reactor/internal/gateway/kubernetes/entity"
)

func ToKubernetesLabelSelector(x entity.LabelSelector) metav1.LabelSelector {
	return metav1.LabelSelector{
		MatchLabels: x.MatchLabels,
		MatchExpressions: ToKubernetesLabelSelectorRequirementList(
			x.MatchExpressions,
		),
	}
}

func ToKubernetesLabelSelectorPtr(
	x *entity.LabelSelector,
) *metav1.LabelSelector {
	if x == nil {
		return nil
	}

	labelSelector := ToKubernetesLabelSelector(*x)

	return &labelSelector
}
