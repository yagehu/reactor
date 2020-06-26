package mapper

import (
	v1 "k8s.io/api/core/v1"

	"github.com/yagehu/reactor/internal/gateway/kubernetes/entity"
)

func ToKubernetesObjectFieldSelector(
	x entity.ObjectFieldSelector,
) v1.ObjectFieldSelector {
	return v1.ObjectFieldSelector{
		APIVersion: x.APIVersion,
		FieldPath:  x.FieldPath,
	}
}

func ToKubernetesObjectFieldSelectorPtr(
	x *entity.ObjectFieldSelector,
) *v1.ObjectFieldSelector {
	if x == nil {
		return nil
	}

	objectFieldSelector := ToKubernetesObjectFieldSelector(*x)

	return &objectFieldSelector
}
