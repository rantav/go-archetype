package transformer

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransformerTransform(t *testing.T) {
	tmp := ".tmp/yyy"
	os.RemoveAll(tmp)
	ts := Transformations{}
	err := Transform(".", tmp, ts)

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
