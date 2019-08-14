package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEvaluateCondition(t *testing.T) {
	assert := assert.New(t)
	tester := func(name string, expected bool, shouldErr bool, condition string, vars map[string]string) {
		truthy, err := EvaluateCondition(condition, vars)
		if shouldErr {
			assert.Errorf(err, "Failed, expected error but there is none: %s", name)
			return
		}
		require.NoErrorf(t, err, "Failed, error not expected: %s", name)
		assert.Equalf(expected, truthy, "Failed, truthy value incorrect in test case: %s", name)
	}
	tests := []struct {
		name      string
		expected  bool
		shouldErr bool
		condition string
		vars      map[string]string
	}{
		{
			"simple condition, truthy value set",
			true,
			false,
			"simpleCondition",
			map[string]string{"simpleCondition": "true"},
		},
		{
			"simple condition, truthy value not set",
			false,
			false,
			"simpleCondition",
			map[string]string{"simpleCondition": ""},
		},
		{
			"condition with and, both are truthy",
			true,
			false,
			`and .c1 .c2`,
			map[string]string{"c1": "true", "c2": "true"},
		},
		{
			"condition with and, one is falsy",
			false,
			false,
			`and .c1 .c2`,
			map[string]string{"c1": "true", "c2": ""},
		},
		{
			"invalid syntax (no dot)",
			false,
			true,
			`and c1 c2`,
			map[string]string{"c1": "", "c2": ""},
		},
	}

	for _, test := range tests {
		tester(test.name, test.expected, test.shouldErr, test.condition, test.vars)
	}
}

func TestEvaluateConditionSad(t *testing.T) {
	assert := assert.New(t)
	_, err := EvaluateCondition("!.x", nil)
	assert.Error(err)
}
