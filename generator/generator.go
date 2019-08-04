package generator

import (
	"github.com/davecgh/go-spew/spew"

	"github.com/rantav/go-archetype/inputs"
	"github.com/rantav/go-archetype/transformer"
	"github.com/rantav/go-archetype/types"
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

	spew.Dump(transformations)
	sourcePath := types.Path(source)
	destinationPath := types.Path(destination)
	err = transformer.Transform(sourcePath, destinationPath, *transformations)
	if err != nil {
		return err
	}

	return nil
}
