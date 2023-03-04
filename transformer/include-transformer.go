package transformer

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/rantav/go-archetype/log"
	"github.com/rantav/go-archetype/template"
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

	logger log.Logger
}

func newIncludeTransformer(spec transformationSpec, logger log.Logger) *includeTransformer {
	return &includeTransformer{
		name:         spec.Name,
		condition:    spec.Condition,
		regionMarker: spec.RegionMarker,
		files:        types.NewFilePatterns(spec.Files),
		logger:       logger,
	}
}

func (t *includeTransformer) GetName() string {
	return t.name
}

func (t *includeTransformer) GetFilePatterns() []types.FilePattern {
	return t.files
}

func (t *includeTransformer) Transform(input types.File) types.File {
	if len(t.regionMarker) == 0 {
		if t.truthy {
			return input
		}
		// Discard the entire file
		return types.File{
			Discarded:    true,
			FullPath:     input.FullPath,
			RelativePath: input.RelativePath,
		}
	}
	// Locate begin and end lines of the markers
	beginMarker := fmt.Sprintf("BEGIN %s", t.regionMarker)
	endMarker := fmt.Sprintf("END %s", t.regionMarker)
	reader := bufio.NewReader(strings.NewReader(input.Contents))
	var (
		output      strings.Builder
		insideBlock = false
	)
	for {
		line, err := read(reader)
		if err != nil {
			if err != io.EOF {
				t.logger.Errorf("Error while scanning file %s: %+v.\n\n Contents: %s ...",
					err, input.FullPath, input.Contents[:100])
			}
			break
		}
		includeLine := t.truthy || !insideBlock
		text := string(line)
		if strings.Contains(text, beginMarker) {
			insideBlock = true
			includeLine = false
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

	newContents := output.String()
	// Check if the last newline should be preserved or discarded.
	if len(newContents) > 0 && !t.hasEmptyLineAtTheEnd(input.Contents) {
		newContents = newContents[:len(newContents)-1]
	}

	return types.File{
		Contents:     newContents,
		FullPath:     input.FullPath,
		RelativePath: input.RelativePath,
	}
}

func (t *includeTransformer) Template(vars map[string]string) error {
	var err error
	t.truthy, err = template.EvaluateCondition(t.condition, vars)
	return err
}

func (t *includeTransformer) hasEmptyLineAtTheEnd(s string) bool {
	l := len(s)
	return l >= 1 && s[l-1] == '\n'
}

// read recursively reads a line into a []byte array,
// supporting lines longer than an arbitrary buffer size
// (which results in "token too long" errors)
//
// Credit: https://devmarkpro.com/working-big-files-golang#approach2
func read(r *bufio.Reader) ([]byte, error) {
	var (
		isPrefix = true
		err      error
		line, ln []byte
	)

	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}

	return ln, err
}

var _ Transformer = &includeTransformer{}
