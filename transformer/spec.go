package transformer

import (
	"github.com/rantav/go-archetype/inputs"
	"github.com/rantav/go-archetype/types"
)

const (
	TransformationTypeInclude = "include"
	TransformationTypeReplace = "replace"
)

type transformationsSpec struct {
	Ignore          []types.FilePattern  `yaml:"ignore"`
	Inputs          []inputs.InputSpec   `yaml:"inputs"`
	Transformations []transformationSpec `yaml:"transformations"`
}

type transformationSpec struct {
	Name         string              `yaml:"name"`
	Type         string              `yaml:"type"`
	Pattern      string              `yaml:"pattern"`
	Replacement  string              `yaml:"replacement"`
	Files        []types.FilePattern `yaml:"files"`
	Condition    string              `yaml:"condition"`
	RegionMarker string              `yaml:"region_marker"`
}
