package transformer

import (
	"os"
	"testing"

	"github.com/rantav/go-archetype/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransformerTransform(t *testing.T) {
	tmp := ".tmp/yyy"
	os.RemoveAll(tmp)
	ts := Transformations{}
	err := Transform(".", tmp, ts, log.NopLogger{})

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
