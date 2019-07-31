package reader

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"

	"github.com/rantav/go-archetype/types"
)

func ReadFile(path types.Path, info os.FileInfo, isIgnored func(types.Path) bool) (
	isDir, ignored bool, contents types.FileContents, err error,
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
	contentsBytes, err := ioutil.ReadFile(string(path))
	if err != nil {
		err = errors.Wrap(err, "reading file")
		return
	}
	contents = types.FileContents(contentsBytes)
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
