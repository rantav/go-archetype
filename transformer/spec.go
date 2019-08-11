package transformer

import (
	"github.com/rantav/go-archetype/inputs"
)

const (
	TransformationTypeInclude = "include"
	TransformationTypeReplace = "replace"
)

type transformationsSpec struct {
	Ignore          []string             `yaml:"ignore"`
	Inputs          []inputs.InputSpec   `yaml:"inputs"`
	Transformations []transformationSpec `yaml:"transformations"`
}

type transformationSpec struct {
	Name         string   `yaml:"name"`
	Type         string   `yaml:"type"`
	Pattern      string   `yaml:"pattern"`
	Replacement  string   `yaml:"replacement"`
	Files        []string `yaml:"files"`
	Condition    string   `yaml:"condition"`
	RegionMarker string   `yaml:"region_marker"`
}
