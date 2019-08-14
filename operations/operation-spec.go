package operations

type Spec struct {
	Operations []OperationSpec `yaml:"operations"`
}
type OperationSpec struct {
	Sh []string `yaml:"sh"`
}

func FromSpec(specs Spec) []Operator {
	var operators []Operator
	for _, s := range specs.Operations {
		op := NewOperator(s)
		operators = append(operators, op)
	}
	return operators
}
