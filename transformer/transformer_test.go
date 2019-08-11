package transformer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransformerTransform(t *testing.T) {
	ts := Transformations{}
	err := Transform(".", ".tmp/yyy", ts)

	require.NoError(t, err)

}
