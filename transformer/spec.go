package transformer

import (
	"github.com/rantav/go-archetype/inputs"
	"github.com/rantav/go-archetype/types"
)

type transformationsSpec struct {
	Ignore          []types.FilePattern  `yaml:"ignore"`
	Inputs          []inputs.InputSpec   `yaml:"inputs"`
	Transformations []transformationSpec `yaml:"transformations"`
}

type transformationSpec struct {
	Name        string              `yaml:"name"`
	Pattern     string              `yaml:"pattern"`
	Replacement string              `yaml:"replacement"`
	Files       []types.FilePattern `yaml:"files"`
}
