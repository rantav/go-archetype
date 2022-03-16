package writer

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/rantav/go-archetype/log"
	"github.com/rantav/go-archetype/types"
)

func WriteFile(destinationBase string, file types.File, mode os.FileMode) error {
	if file.Discarded {
		log.Debugf("File is discarded, not writing: %s", file.RelativePath)
		return nil
	}
	destinationPath := filepath.Join(destinationBase, file.RelativePath)
	log.Infof("Writing file %s", destinationPath)
	dir := filepath.Dir(destinationPath)
	err := os.MkdirAll(dir, os.ModeDir|os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating base dir for file: %w", err)
	}
	if err = ioutil.WriteFile(destinationPath, []byte(file.Contents), mode); err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}
	return nil
}
