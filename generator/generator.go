package generator

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/rantav/go-archetype/inputs"
	"github.com/rantav/go-archetype/log"
	"github.com/rantav/go-archetype/transformer"
)

// Generate is the main entry point for code generation/transformations.
func Generate(transformationsFile, source, destination string, inputArgs []string, logger log.Logger) error {
	transformations, err := transformer.Read(transformationsFile, logger)
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

	vars := collectSystemAndEnvironmentVariables(source, destination)
	err = transformations.Template(vars)
	if err != nil {
		return err
	}

	err = transformer.Transform(source, destination, *transformations, logger)
	if err != nil {
		return err
	}

	return nil
}

// Collects environment variables as well as system variables, e.g. source and destination
func collectSystemAndEnvironmentVariables(source, destination string) map[string]string {
	vars := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		vars[pair[0]] = vars[pair[1]]
	}
	vars["source"] = source
	vars["destination"] = destination
	vars["source_dirname"] = filepath.Base(source)
	vars["destination_dirname"] = filepath.Base(destination)
	return vars
}
