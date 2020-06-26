package mapper

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/yagehu/reactor/internal/gateway/kubernetes/entity"
)

func ToKubernetesTypeMeta(x entity.TypeMeta) metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       x.Kind,
		APIVersion: x.APIVersion,
	}
}
