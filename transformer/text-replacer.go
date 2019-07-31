package transformer

import (
	"strings"

	"github.com/rantav/go-archetype/types"
)

type textReplacer struct {
	name        string
	pattern     string
	replacement string
	files       []types.FilePattern
}

func (t textReplacer) GetName() string {
	return t.name
}

func (t textReplacer) GetFilePatterns() []types.FilePattern {
	return t.files
}

func (t textReplacer) Transform(input types.FileContents) types.FileContents {
	return types.FileContents(strings.ReplaceAll(string(input), t.pattern, t.replacement))
}

func newTextReplacer(raw rawTransformation) textReplacer {
	return textReplacer{
		name:        raw.Name,
		pattern:     raw.Pattern,
		replacement: raw.Replacement,
		files:       raw.Files,
	}
}
