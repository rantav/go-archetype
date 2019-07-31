package transformer

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func Read(transformationsFile string) (*Transformations, error) {
	yamlFile, err := ioutil.ReadFile(transformationsFile)
	if err != nil {
		return nil, err
	}
	var raw rawTransformations
	err = yaml.Unmarshal(yamlFile, &raw)
	if err != nil {
		return nil, err
	}
	fmt.Println(raw)
	return fromRaw(raw)
}

func fromRaw(raw rawTransformations) (*Transformations, error) {
	transformations := Transformations{
		Ignore: raw.Ignore,
	}
	for _, t := range raw.Transformations {
		transformer, err := newTransformer(t)
		if err != nil {
			return nil, err
		}
		transformations.Transformers = append(transformations.Transformers, transformer)
	}
	return &transformations, nil
}

func newTransformer(raw rawTransformation) (Transformer, error) {
	return newTextReplacer(raw), nil
}
