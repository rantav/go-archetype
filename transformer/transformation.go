package transformer

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/rantav/go-archetype/types"
)

type Transformer interface {
	GetName() string
	GetFilePatterns() []types.FilePattern
	Transform(types.FileContents) types.FileContents
}

type Transformations struct {
	Ignore       []types.FilePattern
	Transformers []Transformer
}

type rawTransformations struct {
	Ignore          []types.FilePattern `yaml:"ignore"`
	Transformations []rawTransformation `yaml:"transformations"`
}

type rawTransformation struct {
	Name        string              `yaml:"name"`
	Pattern     string              `yaml:"pattern"`
	Replacement string              `yaml:"replacement"`
	Files       []types.FilePattern `yaml:"files"`
}

func (transformations Transformations) Transform(name types.Path, contents types.FileContents) (
	newContenets types.FileContents, fileIgnored bool, err error,
) {
	if transformations.globallyIgnored(name) {
		return contents, true, nil
	}
	for _, transformer := range transformations.Transformers {
		if ignored(name, transformer.GetFilePatterns()) {
			continue
		}
		contents = transformer.Transform(contents)
	}
	return contents, false, nil
}

func (transformations Transformations) globallyIgnored(name types.Path) bool {
	return ignored(name, transformations.Ignore)
}

func ignored(name types.Path, patterns []types.FilePattern) bool {
	for _, pattern := range patterns {
		// Check glob match
		matched, err := filepath.Match(string(pattern), string(name))
		if err != nil {
			log.Printf("Error matching pattern %s to file %s. %+v \n", pattern, name, err)
		}
		if matched {
			return true
		}

		// And check string prefix match (when / is used at the end)
		if strings.HasSuffix(string(pattern), "/") {
			if strings.HasPrefix(string(name), string(pattern)) {
				return true
			}
		}
	}
	return false
}
