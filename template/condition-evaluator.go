package template

import (
	"bytes"
	"fmt"
	"strings"
	tt "text/template"
)

// Evaluates the condition in the syntax of go template if conditions and returns true
// if this condition holds true.
func EvaluateCondition(condition string, vars map[string]string) (bool, error) {
	// Create a template and contains the condition in order to evaluate it
	condition = fixSingleTermConditionSpecialCase(condition)
	text := fmt.Sprintf("{{if %s}}true{{end}}", condition)
	tmpl, err := tt.New("t").Parse(text)
	if err != nil {
		return false, fmt.Errorf("error creating the text template: %w", err)
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, vars)
	if err != nil {
		return false, fmt.Errorf("error templating: %w", err)
	}
	val := buf.String() == "true"
	return val, nil
}

func fixSingleTermConditionSpecialCase(condition string) string {
	if isSingleWord(condition) && !startWithDot(condition) {
		return fmt.Sprintf(".%s", condition)
	}
	return condition
}

func isSingleWord(str string) bool {
	return !strings.Contains(str, " ")
}

func startWithDot(str string) bool {
	return len(str) > 0 && str[0] == '.'
}
