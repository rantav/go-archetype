package operations

type Spec struct {
	Operations []OperationSpec `yaml:"operations"`
}
type OperationSpec struct {
	Sh []shellOperationSpec `yaml:"sh"`
}

type shellOperationSpec struct {
	Cmd       string `yaml:"cmd"`
	Multiline bool   `yaml:"multiline"`
}

func (s *shellOperationSpec) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var cmd string
	if err := unmarshal(&cmd); err == nil {
		s.Cmd = cmd
		return nil
	}

	type rawShellOperation shellOperationSpec
	if err := unmarshal((*rawShellOperation)(s)); err != nil {
		return err
	}

	return nil
}

func FromSpec(specs Spec) []Operator {
	var operators []Operator
	for _, s := range specs.Operations {
		op := NewOperator(s)
		operators = append(operators, op)
	}
	return operators
}
