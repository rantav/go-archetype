package operations

import (
	"github.com/rantav/go-archetype/log"
)

type Operator interface {
	Operate() error
	Template(vars map[string]string) error
}

func NewOperator(spec OperationSpec, logger log.Logger) Operator {
	return newShellOperator(spec, logger) // Right now it's the only supported operator
}
