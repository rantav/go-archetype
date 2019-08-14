package generator

import (
	"github.com/davecgh/go-spew/spew"

	"github.com/rantav/go-archetype/inputs"
	"github.com/rantav/go-archetype/log"
	"github.com/rantav/go-archetype/transformer"
)

// Generate is the main entry point for code generation/transformations.
func Generate(transformationsFile, source, destination string, inputArgs []string) error {
	transformations, err := transformer.Read(transformationsFile)
	if err != nil {
		return err
	}

	err = inputs.ParseCLIArgsInputs(transformations, inputArgs)
	if err != nil {
		return err
	}

	err = inputs.CollectUserInputs(transformations)
	if err != nil {
		return err
	}

	err = transformations.Template()
	if err != nil {
		return err
	}

	log.Debugf(spew.Sdump(transformations))
	err = transformer.Transform(source, destination, *transformations)
	if err != nil {
		return err
	}

	return nil
}
