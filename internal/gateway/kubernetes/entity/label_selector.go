package entity

type LabelSelector struct {
	MatchLabels      map[string]string
	MatchExpressions []LabelSelectorRequirement
}

func NewLabelSelector(
	matchLabels map[string]string,
	matchExpressions []LabelSelectorRequirement,
) LabelSelector {
	return LabelSelector{
		MatchLabels:      matchLabels,
		MatchExpressions: matchExpressions,
	}
}

func LabelSelectorPtr(x LabelSelector) *LabelSelector {
	return &x
}
