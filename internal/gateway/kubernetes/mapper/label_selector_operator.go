package mapper

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/yagehu/reactor/internal/gateway/kubernetes/entity"
)

func ToKubernetesLabelSelectorOperator(
	x entity.LabelSelectorOperator,
) metav1.LabelSelectorOperator {
	return metav1.LabelSelectorOperator(x.String())
}
