package transformer

import (
	"testing"

	"github.com/rantav/go-archetype/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTextReplacerBasics(t *testing.T) {
	assert := assert.New(t)

	r := newTextReplacer(transformationSpec{
		Name: "1",
	})
	assert.Equal("1", r.GetName())
	assert.Empty(r.GetFilePatterns())
}

func TestTextReplacerTemplate(t *testing.T) {
	assert := assert.New(t)

	r := newTextReplacer(transformationSpec{
		Replacement: "",
	})
	err := r.Template(nil)
	require.NoError(t, err)
	assert.Equal("", r.replacement)

	r = newTextReplacer(transformationSpec{
		Replacement: "{{ .x }} etc",
	})
	err = r.Template(map[string]string{"x": "X"})
	require.NoError(t, err)
	assert.Equal("X etc", r.replacement)
}

func TestTextReplacerTransform(t *testing.T) {
	assert := assert.New(t)

	r := newTextReplacer(transformationSpec{
		Pattern:     "123",
		Replacement: "456",
	})
	file := r.Transform(types.File{
		Contents: "123",
	})
	assert.Equal(types.File{
		Contents: "456",
	}, file)
}
