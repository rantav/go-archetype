package types

import (
	"fmt"

	"github.com/gobwas/glob"
)

type FilePattern struct {
	Pattern  string
	glob     glob.Glob
	compiled bool
}

func (f *FilePattern) Match(path string) (bool, error) {
	if !f.compiled {
		// Compile once, on demand
		var err error
		f.glob, err = glob.Compile(f.Pattern, '/')
		if err != nil {
			return false, fmt.Errorf("error compiling pattern %s: %w", f.Pattern, err)
		}
		f.compiled = true
	}
	return f.glob.Match(path), nil
}

func NewFilePatterns(paths []string) []FilePattern {
	var patterns []FilePattern
	for _, p := range paths {
		patterns = append(patterns, FilePattern{Pattern: p})
	}
	return patterns
}

type File struct {
	Contents string
	// The full path to the original (source) file
	FullPath string
	// The relative path to the file (relative to the root of the project, e.g. source)
	RelativePath string
	// Mark this file as needs to be discarded (as opposed to just an empty file)
	Discarded bool
}
