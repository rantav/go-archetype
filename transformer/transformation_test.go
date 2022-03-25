package transformer

import (
	"testing"

	"github.com/rantav/go-archetype/inputs"
	"github.com/rantav/go-archetype/log"
	"github.com/rantav/go-archetype/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransformationsTransform(t *testing.T) {
	assert := assert.New(t)

	// empty transformers list
	ts := Transformations{}
	file, err := ts.Transform(types.File{
		Contents: "x",
	})
	require.NoError(t, err)
	assert.Equal("x", file.Contents)

	// Single replacer
	ts = Transformations{
		transformers: []Transformer{newTextReplacer(
			transformationSpec{
				Pattern:     "x",
				Replacement: "y",
				Files:       []string{"*.go"},
			})},
		logger: log.NopLogger{},
	}
	file, err = ts.Transform(types.File{
		Contents:     "x",
		RelativePath: "hello.go",
	})
	require.NoError(t, err)
	assert.Equal("y", file.Contents)

	// A file that doesn't match
	ts = Transformations{
		transformers: []Transformer{newTextReplacer(
			transformationSpec{
				Pattern:     "x",
				Replacement: "y",
				Files:       []string{"hello.go"},
			})},
	}
	file, err = ts.Transform(types.File{
		RelativePath: "go.away",
		Contents:     "x",
	})
	require.NoError(t, err)
	assert.Equal("x", file.Contents)
}

func TestTransformationsTemplate(t *testing.T) {
	// empty transformers list
	ts := Transformations{
		transformers: []Transformer{
			newTextReplacer(transformationSpec{}),
		},
		prompters: []inputs.Prompter{
			inputs.NewPrompt(inputs.InputSpec{Type: "text"}),
		},
		logger: log.NopLogger{},
	}
	err := ts.Template(make(map[string]string))
	require.NoError(t, err)
}

func TestTransformationsMatched(t *testing.T) {
	assert := assert.New(t)

	tr := Transformations{logger: log.NopLogger{}}

	assert.True(tr.matched("hello.go", []types.FilePattern{{Pattern: "hello.go"}}, false))
	assert.True(tr.matched("all/hello.go", []types.FilePattern{{Pattern: "all/"}}, true))

	// Test some globs
	assert.True(tr.matched("hello.go", []types.FilePattern{{Pattern: "*.go"}}, false))
	assert.True(tr.matched("x/hello.go", []types.FilePattern{{Pattern: "*/*.go"}}, false))
	assert.False(tr.matched("x/hello.go", []types.FilePattern{{Pattern: "*.go"}}, false))
	assert.True(tr.matched("x/y/hello.go", []types.FilePattern{{Pattern: "**/*.go"}}, false))
}
