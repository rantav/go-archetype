package transformer

import (
	"testing"

	"github.com/rantav/go-archetype/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {
	assert := assert.New(t)

	transformations, err := Read("./testdata/transformations.yml", log.NopLogger{})
	require.NoError(t, err)
	assert.NotNil(transformations)

	assert.Len(transformations.GetInputPrompters(), 3)
	assert.Len(transformations.transformers, 5)
}
