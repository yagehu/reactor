package mapper

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/yagehu/reactor/internal/gateway/kubernetes/entity"
)

func ToKubernetesObjectMeta(x entity.ObjectMeta) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Name:      x.Name,
		Namespace: x.Namespace,
		Labels:    x.Labels,
	}
}
