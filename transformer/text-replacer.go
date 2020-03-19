package transformer

import (
	"strings"

	"github.com/rantav/go-archetype/template"
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

func (t *textReplacer) Transform(input types.File) types.File {
	return types.File{
		Contents:     strings.ReplaceAll(input.Contents, t.pattern, t.replacement),
		FullPath:     input.FullPath,
		RelativePath: input.RelativePath,
		Discarded:    input.Discarded,
	}
}

func (t *textReplacer) Template(vars map[string]string) error {
	var err error
	t.replacement, err = template.Execute(t.replacement, vars)
	return err
}

func newTextReplacer(spec transformationSpec) *textReplacer {
	return &textReplacer{
		name:        spec.Name,
		pattern:     spec.Pattern,
		replacement: spec.Replacement,
		files:       types.NewFilePatterns(spec.Files),
	}
}

var _ Transformer = &textReplacer{}
