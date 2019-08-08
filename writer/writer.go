package writer

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/rantav/go-archetype/log"
	"github.com/rantav/go-archetype/types"
)

func WriteFile(destinationBase, file types.Path, contents types.FileContents, mode os.FileMode) error {
	destinationPath := filepath.Join(string(destinationBase), string(file))
	log.Infof("Writing file %s", destinationPath)
	dir := filepath.Dir(destinationPath)
	err := os.MkdirAll(dir, os.ModeDir|os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "error creating base dir for file")
	}
	err = ioutil.WriteFile(destinationPath, []byte(contents), mode)
	return errors.Wrap(err, "error writing file")
}
