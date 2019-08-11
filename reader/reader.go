package reader

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"

	"github.com/rantav/go-archetype/types"
)

func ReadFile(path string, info os.FileInfo, isIgnored func(string) bool) (
	isDir, ignored bool, file types.File, err error,
) {
	isDir, err = isDirectory(info)
	if err != nil {
		err = errors.Wrap(err, "checking if dir")
		return
	}
	if isDir {
		return
	}
	ignored = isIgnored(path)
	if ignored {
		return
	}
	contentsBytes, err := ioutil.ReadFile(path)
	if err != nil {
		err = errors.Wrap(err, "reading file")
		return
	}
	file = types.File{
		Contents: string(contentsBytes),
		Path:     path,
	}
	return
}

func isDirectory(fi os.FileInfo) (bool, error) {
	switch mode := fi.Mode(); {
	case mode.IsDir():
		return true, nil
	case mode.IsRegular():
		return false, nil
	default:
		return false, fmt.Errorf("unknown file mode (dir or file) at %s", fi)
	}
}
