package mapper

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/yagehu/reactor/internal/gateway/kubernetes/entity"
)

func ToKubernetesLabelSelectorRequirement(
	x entity.LabelSelectorRequirement,
) metav1.LabelSelectorRequirement {
	return metav1.LabelSelectorRequirement{
		Key:      x.Key,
		Operator: ToKubernetesLabelSelectorOperator(x.Operator),
		Values:   x.Values,
	}
}

func ToKubernetesLabelSelectorRequirementList(
	xs []entity.LabelSelectorRequirement,
) []metav1.LabelSelectorRequirement {
	s := make([]metav1.LabelSelectorRequirement, 0, len(xs))

	for _, x := range xs {
		s = append(s, ToKubernetesLabelSelectorRequirement(x))
	}

	return s
}
