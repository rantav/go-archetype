package writer

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/rantav/go-archetype/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const tmpDir = ".tmp/test/"

func TestWriteFileNoDiscard(t *testing.T) {
	cleanup()
	assert := assert.New(t)
	destinationBase := tmpDir
	file := types.File{
		Contents:     "123",
		RelativePath: "file.txt",
		Discarded:    false,
	}
	mode := os.ModePerm
	err := WriteFile(destinationBase, file, mode)
	require.NoError(t, err)

	target := destinationBase + "file.txt"
	assert.FileExists(target)
	stat, err := os.Lstat(target)
	require.NoError(t, err)
	assert.False(stat.IsDir())
	contents, err := ioutil.ReadFile(target)
	require.NoError(t, err)
	assert.Equal("123", string(contents))
}

func TestWriteFileDiscard(t *testing.T) {
	cleanup()
	assert := assert.New(t)
	destinationBase := tmpDir
	file := types.File{
		Contents:     "123",
		RelativePath: "discarded",
		Discarded:    true,
	}
	mode := os.ModePerm
	err := WriteFile(destinationBase, file, mode)
	require.NoError(t, err)

	target := destinationBase + "discarded"
	_, err = os.Lstat(target)
	assert.Error(err)
	assert.Contains(err.Error(), "no such file or directory")
}

func cleanup() {
	err := os.RemoveAll(tmpDir)
	if err != nil {
		panic(err)
	}
}
