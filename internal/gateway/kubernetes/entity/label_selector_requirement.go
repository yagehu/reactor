package entity

type LabelSelectorRequirement struct {
	Key      string
	Operator LabelSelectorOperator
	Values   []string
}
