package transformer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {
	assert := assert.New(t)

	transformations, err := Read("../transformations.yml")
	require.NoError(t, err)
	assert.NotNil(transformations)

	assert.Equal(2, len(transformations.GetInputPrompters()))
	assert.Equal(4, len(transformations.transformers))
}
