package transformer

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rantav/go-archetype/log"
	"github.com/rantav/go-archetype/types"
)

func TestTransform(t *testing.T) {
	assert := assert.New(t)
	tester := func(name string, truthy bool, marker string, input, expectedOutput string, expectedDiscarded bool) {
		transformer := includeTransformer{
			truthy:       truthy,
			regionMarker: marker,
			logger:       log.NopLogger{},
		}
		output := transformer.Transform(types.File{Contents: input})
		assert.Equalf(expectedOutput, output.Contents, "Test failed, output not expected: %s", name)
		assert.Equalf(expectedDiscarded, output.Discarded, "Test failed, discarded not expected: %s", name)
	}

	// make a string just a bit longer than the scanner max
	repeatMe := "I will not waste chalk. "
	excessivelyLongLine := strings.Repeat(repeatMe, 1+bufio.MaxScanTokenSize/len(repeatMe))

	// nolint:maligned
	tests := []struct {
		name      string
		truthy    bool
		marker    string
		input     string
		output    string
		discarded bool
	}{
		{
			"truthy, all empty",
			true,
			"",
			"",
			"",
			false,
		},
		{
			"falsy, all empty",
			false,
			"",
			"",
			"",
			true,
		},
		{
			"truthy, with a marker",
			true,
			"__1__",
			`1
BEGIN __1__
2
END __1__
3`,
			`1
2
3`,
			false,
		},
		{
			"falsy, with a marker",
			false,
			"__1__",
			`1
BEGIN __1__
2
END __1__
3
`,
			`1
3
`,
			false,
		},
		{
			"truthy, non-empty file, but no marker",
			true,
			"",
			"1",
			"1",
			false,
		},
		{
			"falsy, non-empty file, but no marker",
			false,
			"",
			"1",
			"",
			true,
		},
		{
			"falsy, with a marker",
			false,
			"__1__",
			`1
BEGIN __1__
2
END __1__
`,
			`1
`,
			false,
		},
		{
			"excessively long line",
			false,
			"__1__",
			excessivelyLongLine,
			excessivelyLongLine,
			false,
		},
	}
	for _, test := range tests {
		tester(test.name, test.truthy, test.marker, test.input, test.output, test.discarded)
	}
}

func TestIncludeTransformerTemplate(t *testing.T) {
	assert := assert.New(t)
	transformer := includeTransformer{}
	vars := make(map[string]string)
	err := transformer.Template(vars)
	assert.NoError(err)
}

func TestIncludeTransformerBasics(t *testing.T) {
	assert := assert.New(t)

	r := newIncludeTransformer(transformationSpec{
		Name: "1",
	}, log.NopLogger{})
	assert.Equal("1", r.GetName())
	assert.Empty(r.GetFilePatterns())
}
