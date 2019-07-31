package transformer

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/pkg/errors"

	"github.com/rantav/go-archetype/reader"
	"github.com/rantav/go-archetype/types"
	"github.com/rantav/go-archetype/writer"
)

func Transform(source, destination types.Path, transformations Transformations) error {
	return filepath.Walk(string(source), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrap(err, "error walking to file")
		}
		sourceFile := types.Path(path)
		isDir, ignored, contents, err := reader.ReadFile(sourceFile, info, transformations.IsGloballyIgnored)
		if err != nil {
			return errors.Wrap(err, "error reading file")
		}
		if isDir {
			return nil
		}

		contents, err = transformations.Transform(sourceFile, contents)
		if ignored {
			color.Blue("Ignoring file %s", path)
		} else {
			err := writer.WriteFile(destination, sourceFile, contents, info.Mode())
			if err != nil {
				return err
			}
		}
		return errors.Wrap(err, "transforming")
	})
}
