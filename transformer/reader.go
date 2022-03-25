package transformer

import (
	"io/ioutil"

	"github.com/rantav/go-archetype/log"
	"gopkg.in/yaml.v2"

	"github.com/rantav/go-archetype/inputs"
	"github.com/rantav/go-archetype/operations"
	"github.com/rantav/go-archetype/types"
)

func Read(transformationsFile string, logger log.Logger) (*Transformations, error) {
	yamlFile, err := ioutil.ReadFile(transformationsFile)
	if err != nil {
		return nil, err
	}
	var spec transformationsSpec
	err = yaml.Unmarshal(yamlFile, &spec)
	if err != nil {
		return nil, err
	}
	return FromSpec(spec, logger)
}

func FromSpec(spec transformationsSpec, logger log.Logger) (*Transformations, error) {
	return &Transformations{
		ignore:       types.NewFilePatterns(spec.Ignore),
		transformers: transformersFromSpec(spec.Transformations, logger),
		prompters:    inputs.FromSpec(spec.Inputs),
		userInputs:   make(map[string]inputs.PromptResponse),
		before:       operations.FromSpec(spec.Before, logger),
		after:        operations.FromSpec(spec.After, logger),

		logger: logger,
	}, nil
}

func transformersFromSpec(transformationSpecs []transformationSpec, logger log.Logger) []Transformer {
	var transformers []Transformer
	for _, t := range transformationSpecs {
		transformers = append(transformers, newTransformer(t, logger))
	}
	return transformers
}
