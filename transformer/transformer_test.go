package transformer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransformerTransform(t *testing.T) {
	ts := Transformations{}
	err := Transform(".", ".tmp/yyy", ts)

	require.NoError(t, err)
}

func TestIsDirEmptyOrDoesntExist(t *testing.T) {
	empty, err := isDirEmptyOrDoesntExist(".")
	assert.NoError(t, err)
	assert.False(t, empty)

	empty, err = isDirEmptyOrDoesntExist(".does-not-exist")
	assert.NoError(t, err)
	assert.True(t, empty)
}
