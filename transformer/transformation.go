package transformer

import (
	"strings"

	"github.com/rantav/go-archetype/inputs"
	"github.com/rantav/go-archetype/log"
	"github.com/rantav/go-archetype/operations"
	"github.com/rantav/go-archetype/types"
)

type Transformations struct {
	// Global ignore patterns
	ignore []types.FilePattern

	// The list of transformers
	transformers []Transformer

	// User prompters
	prompters []inputs.Prompter

	// User's responses to prompters
	userInputs map[string]inputs.PromptResponse

	// Before and after operators
	before []operations.Operator
	after  []operations.Operator
}

func (t Transformations) Transform(file types.File) (
	newFile types.File, err error,
) {
	for _, transformer := range t.transformers {
		// See if file already needs to be discarded
		if file.Discarded {
			return file, nil
		}
		if !matched(file.RelativePath, transformer.GetFilePatterns(), false) {
			continue
		}
		log.Debugf("Applying transformer [%s] to file [%s]", transformer.GetName(), file.RelativePath)
		file = transformer.Transform(file)
	}
	return file, nil
}

func (t Transformations) IsGloballyIgnored(path string) bool {
	return matched(path, t.ignore, true)
}

func (t Transformations) GetInputPrompters() []inputs.Prompter {
	return t.prompters
}

func (t *Transformations) SetResponse(response inputs.PromptResponse) {
	t.userInputs[response.ID] = response
}

func (t *Transformations) Template(vars map[string]string) error {
	// Add user inputs
	for _, input := range t.userInputs {
		vars[input.ID] = input.Answer
	}

	for _, tr := range t.transformers {
		err := tr.Template(vars)
		if err != nil {
			return err
		}
	}

	for _, op := range t.before {
		err := op.Template(vars)
		if err != nil {
			return err
		}
	}
	for _, op := range t.after {
		err := op.Template(vars)
		if err != nil {
			return err
		}
	}
	return nil
}

func matched(path string, patterns []types.FilePattern, includePrefix bool) bool {
	for _, pattern := range patterns {
		// Check glob match
		matched, err := pattern.Match(path)
		if err != nil {
			log.Errorf("Error matching pattern %s to file %s. %+v \n", pattern.Pattern, path, err)
		}
		if matched {
			return true
		}

		if includePrefix {
			// And check string prefix match (when / is used at the end)
			if strings.HasSuffix(pattern.Pattern, "/") {
				if strings.HasPrefix(path, pattern.Pattern) {
					return true
				}
			}
		}
	}
	return false
}
