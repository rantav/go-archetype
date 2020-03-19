package reader

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	"github.com/rantav/go-archetype/types"
)

func ReadFile(path string, info os.FileInfo, sourceDir string, isIgnored func(string) bool) (
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

	relativePath := relative(sourceDir, path)
	ignored = isIgnored(relativePath)
	if ignored {
		return
	}
	contentsBytes, err := ioutil.ReadFile(path)
	if err != nil {
		err = errors.Wrap(err, "reading file")
		return
	}
	file = types.File{
		Contents:     string(contentsBytes),
		FullPath:     path,
		RelativePath: relativePath,
	}
	return
}

// Create a relative path from path by removing the prefix if necessary.
func relative(prefix, path string) string {
	if filepath.Clean(prefix) == "." {
		// Nothing to remove, empty prefix (or ".")
		return path
	}
	if !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}
	return filepath.Clean(strings.TrimPrefix(path, prefix))
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
