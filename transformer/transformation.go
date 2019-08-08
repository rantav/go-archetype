package transformer

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/rantav/go-archetype/inputs"
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
}

func (t Transformations) Transform(name types.Path, contents types.FileContents) (
	newContenets types.FileContents, err error,
) {
	for _, transformer := range t.transformers {
		if !matched(name, transformer.GetFilePatterns(), false) {
			continue
		}
		contents = transformer.Transform(contents)
	}
	return contents, nil
}

func (t Transformations) IsGloballyIgnored(name types.Path) bool {
	return matched(name, t.ignore, true)
}

func (t Transformations) GetInputPrompters() []inputs.Prompter {
	return t.prompters
}

func (t *Transformations) SetResponse(response inputs.PromptResponse) {
	t.userInputs[response.ID] = response
}

func (t *Transformations) Template() error {
	inputsAsMap := make(map[string]string)
	for _, input := range t.userInputs {
		inputsAsMap[input.ID] = input.Answer
	}
	for _, tr := range t.transformers {
		err := tr.Template(inputsAsMap)
		if err != nil {
			return err
		}
	}
	return nil
}

func matched(name types.Path, patterns []types.FilePattern, includePrefix bool) bool {
	for _, pattern := range patterns {
		// Check glob match
		matched, err := filepath.Match(string(pattern), string(name))
		if err != nil {
			log.Printf("Error matching pattern %s to file %s. %+v \n", pattern, name, err)
		}
		if matched {
			return true
		}

		if includePrefix {
			// And check string prefix match (when / is used at the end)
			if strings.HasSuffix(string(pattern), "/") {
				if strings.HasPrefix(string(name), string(pattern)) {
					return true
				}
			}
		}
	}
	return false
}
