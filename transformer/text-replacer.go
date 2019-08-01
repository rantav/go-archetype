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

func (t *textReplacer) GetName() string {
	return t.name
}

func (t *textReplacer) GetFilePatterns() []types.FilePattern {
	return t.files
}

func (t *textReplacer) Transform(input types.FileContents) types.FileContents {
	return types.FileContents(strings.ReplaceAll(string(input), t.pattern, t.replacement))
}

func (t *textReplacer) Template(vars map[string]string) error {
	var err error
	t.replacement, err = template(t.replacement, vars)
	return err
}

func newTextReplacer(spec transformationSpec) *textReplacer {
	return &textReplacer{
		name:        spec.Name,
		pattern:     spec.Pattern,
		replacement: spec.Replacement,
		files:       spec.Files,
	}
}
