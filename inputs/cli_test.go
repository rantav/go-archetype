package inputs

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseUserInputCliArgs(t *testing.T) {
	assert := assert.New(t)
	sp := func(s string) *string { return &s }

	parsed, err := parseUserInputCliArgs([]string{}, []string{})
	require.NoError(t, err)
	assert.Empty(parsed)

	parsed, err = parseUserInputCliArgs([]string{"hello"}, []string{})
	require.NoError(t, err)
	assert.Equal([]*string{nil}, parsed)

	parsed, err = parseUserInputCliArgs([]string{}, []string{"hello"})
	require.NoError(t, err)
	assert.Empty(parsed)

	parsed, err = parseUserInputCliArgs([]string{"hello"}, []string{"--hello", "world"})
	require.NoError(t, err)
	assert.Equal([]*string{sp("world")}, parsed)

	parsed, err = parseUserInputCliArgs([]string{"hello", "there"}, []string{"--hello", "world"})
	require.NoError(t, err)
	assert.Equal([]*string{sp("world"), nil}, parsed)

	// Mix the order of CLI args
	parsed, err = parseUserInputCliArgs([]string{"hello", "there"}, []string{"--there", "x", "--hello", "world"})
	require.NoError(t, err)
	assert.Equal([]*string{sp("world"), sp("x")}, parsed)
}

func TestParseCLIArgsInputs(t *testing.T) {
	assert := assert.New(t)
	p := newSimpleTextPrompter(InputSpec{
		Type: "text",
		ID:   "t",
	})

	err := ParseCLIArgsInputs(&mockInputsCollector{prompters: []Prompter{p}}, []string{})
	require.NoError(t, err)
	assert.False(p.Answered)

	err = ParseCLIArgsInputs(&mockInputsCollector{prompters: []Prompter{p}}, []string{"--t", "y"})
	require.NoError(t, err)
	assert.True(p.Answered)
	assert.Equal("y", p.Answer)
}
