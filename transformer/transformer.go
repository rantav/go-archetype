package transformer

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/rantav/go-archetype/types"
)

func Transform(source, destination types.Path, transformations Transformations) error {
	return filepath.Walk(string(source), func(path string, info os.FileInfo, err error) error {
		isDir, err := isDirectory(path)
		if err != nil {
			return errors.Wrap(err, "checking if dir")
		}
		if isDir {
			return nil
		}
		sourceFile := types.Path(path)
		contents, err := ioutil.ReadFile(path)
		if err != nil {
			return errors.Wrap(err, "reading file")
		}

		result, ignored, err := transformations.Transform(sourceFile, types.FileContents(contents))
		if !ignored {
			fmt.Println("File", path)
			fmt.Println(result)
		}
		return errors.Wrap(err, "transforming")
	})
}

func isDirectory(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		return true, nil
	case mode.IsRegular():
		return false, nil
	default:
		return false, fmt.Errorf("unknown file mode (dir or file) at %s", err)
	}
}
