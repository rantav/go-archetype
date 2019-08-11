package transformer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTempalte(t *testing.T) {
	assert := assert.New(t)

	result, err := template("", nil)
	require.NoError(t, err)
	assert.Empty(result)

	// buggy template
	_, err = template("{{}}", nil)
	assert.Error(err)

	result, err = template("hello {{.x}}", map[string]string{"x": "world"})
	require.NoError(t, err)
	assert.Equal("hello world", result)
}
