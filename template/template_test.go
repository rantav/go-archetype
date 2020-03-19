package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTempalte(t *testing.T) {
	assert := assert.New(t)

	result, err := Execute("", nil)
	require.NoError(t, err)
	assert.Empty(result)

	// test a buggy template
	_, err = Execute("{{}}", nil)
	assert.Error(err)

	result, err = Execute("hello {{.x}}", map[string]string{"x": "world"})
	require.NoError(t, err)
	assert.Equal("hello world", result)
}
