package operations

type Operator interface {
	Operate() error
	Template(vars map[string]string) error
}

func NewOperator(spec OperationSpec) Operator {
	return newShellOperator(spec) // Right now it's the only supported operator
}
