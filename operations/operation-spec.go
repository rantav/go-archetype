package operations

import (
	"github.com/rantav/go-archetype/log"
)

type Spec struct {
	Operations []OperationSpec `yaml:"operations"`
}
type OperationSpec struct {
	Sh []string `yaml:"sh"`
}

func FromSpec(specs Spec, logger log.Logger) []Operator {
	var operators []Operator
	for _, s := range specs.Operations {
		op := NewOperator(s, logger)
		operators = append(operators, op)
	}
	return operators
}
