package transformer

import (
	"path/filepath"
	"strings"

	"github.com/rantav/go-archetype/template"
	"github.com/rantav/go-archetype/types"
)

type fileRenamer struct {
	name        string
	pattern     string
	replacement string
	files       []types.FilePattern
}

func (t *fileRenamer) GetName() string {
	return t.name
}

func (t *fileRenamer) GetFilePatterns() []types.FilePattern {
	return t.files
}

func (t *fileRenamer) Transform(input types.File) types.File {
	relativePath := strings.ReplaceAll(input.RelativePath, t.pattern, t.replacement)
	fullPath := input.FullPath

	relativeIndex := strings.LastIndex(input.FullPath, filepath.Clean(input.RelativePath))
	if relativeIndex > 0 {
		fullPath = fullPath[:relativeIndex] + filepath.Clean(relativePath)
	}

	return types.File{
		Contents:     input.Contents,
		FullPath:     fullPath,
		RelativePath: relativePath,
		Discarded:    input.Discarded,
	}
}

func (t *fileRenamer) Template(vars map[string]string) error {
	var err error
	t.replacement, err = template.Execute(t.replacement, vars)
	return err
}

func newFileRenamer(spec transformationSpec) *fileRenamer {
	return &fileRenamer{
		name:        spec.Name,
		pattern:     spec.Pattern,
		replacement: spec.Replacement,
		files:       types.NewFilePatterns(spec.Files),
	}
}

var _ Transformer = &fileRenamer{}
