package transformer

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/rantav/go-archetype/log"
	"github.com/rantav/go-archetype/types"
)

type includeTransformer struct {
	name string
	// the condition as a string
	condition string
	// regions marker in the file
	regionMarker string
	files        []types.FilePattern
	// Does this condition evaluate to true, provided the variable values?
	truthy bool
}

func newIncludeTransformer(spec transformationSpec) *includeTransformer {
	return &includeTransformer{
		name:         spec.Name,
		condition:    spec.Condition,
		regionMarker: spec.RegionMarker,
		files:        spec.Files,
	}
}

func (t *includeTransformer) GetName() string {
	return t.name
}

func (t *includeTransformer) GetFilePatterns() []types.FilePattern {
	return t.files
}

func (t *includeTransformer) Transform(input types.FileContents) types.FileContents {
	// Locate begin and end lines of the markers
	beginMarker := fmt.Sprintf("BEGIN %s", t.regionMarker)
	endMarker := fmt.Sprintf("END %s", t.regionMarker)
	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	var (
		output      strings.Builder
		insideBlock = false
		anyFound    = false
	)
	for scanner.Scan() {
		includeLine := t.truthy || !insideBlock
		text := scanner.Text()
		if strings.Contains(text, beginMarker) {
			insideBlock = true
			includeLine = false
			anyFound = true
		}
		if strings.Contains(text, endMarker) {
			insideBlock = false
			includeLine = false
		}
		if includeLine {
			output.WriteString(text)
			output.WriteRune('\n')
		}
	}
	if scanner.Err() != nil {
		log.Errorf("Error while scanning file: %+v.\n\n Contents: %s", scanner.Err(), input)
	}

	if !anyFound && !t.truthy {
		return ""
	}
	return types.FileContents(output.String())
}

func (t *includeTransformer) Template(vars map[string]string) error {
	var err error
	t.truthy, err = evaluateCondition(t.condition, vars)
	return err
}

var _ Transformer = &includeTransformer{}
