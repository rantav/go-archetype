package transformer

import (
	"testing"

	"github.com/rantav/go-archetype/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFileRenamer_Transform(t *testing.T) {
	t.Parallel()

	t.Run("Content should never change", func(t *testing.T) {
		t.Parallel()

		renamer := newFileRenamer(transformationSpec{
			Name:        "project name",
			Type:        "rename",
			Pattern:     "go-arhetype",
			Replacement: "project_name",
			Files:       []string{"*.go"},
		})

		input := types.File{
			Contents:     "Very important content go-archetype",
			FullPath:     "",
			RelativePath: "",
			Discarded:    false,
		}
		expected := types.File{
			Contents:     "Very important content go-archetype",
			FullPath:     "",
			RelativePath: "",
			Discarded:    false,
		}

		actual := renamer.Transform(input)
		assert.Equal(t, expected, actual)
	})

	t.Run("Rename relative path", func(t *testing.T) {
		t.Parallel()

		testCases := []struct {
			name     string
			input    string
			expected string
		}{
			{
				name:     "empty input",
				input:    "",
				expected: "",
			},
			{
				name:     "project name is not present",
				input:    "./very-awesome.go",
				expected: "./very-awesome.go",
			},
			{
				name:     "only project name present",
				input:    "go-archetype",
				expected: "project_name",
			},
			{
				name:     "complex path with directories",
				input:    "./cmd/web/go-archetype",
				expected: "./cmd/web/project_name",
			},
			{
				name:     "complex path with directories end extension",
				input:    "./cmd/web/go-archetype.go",
				expected: "./cmd/web/project_name.go",
			},
			{
				name:     "project name everywhere",
				input:    "./cmd/go-archetype/web/go-archetype.go",
				expected: "./cmd/project_name/web/project_name.go",
			},
		}

		renamer := fileRenamer{
			name:        "project name",
			pattern:     "go-archetype",
			replacement: "project_name",
			files:       nil,
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				t.Parallel()

				input := types.File{
					Contents:     "",
					FullPath:     "",
					RelativePath: testCase.input,
					Discarded:    false,
				}

				actual := renamer.Transform(input)
				assert.Equal(t, testCase.expected, actual.RelativePath)
			})
		}
	})

	t.Run("Rename full path", func(t *testing.T) {
		t.Parallel()

		testCases := []struct {
			name          string
			inputFull     string
			inputRelative string
			expected      string
		}{
			{
				name:          "empty input",
				inputFull:     "",
				inputRelative: "",
				expected:      "",
			},
			{
				name:          "project name is not present",
				inputFull:     "/home/user1/projects/my-project/very-awesome.go",
				inputRelative: "my-project/very-awesome.go",
				expected:      "/home/user1/projects/my-project/very-awesome.go",
			},
			{
				name:          "project name present",
				inputFull:     "/home/user1/projects/my-project/go-archetype",
				inputRelative: "my-project/go-archetype",
				expected:      "/home/user1/projects/my-project/project_name",
			},
			{
				name:          "complex path with directories end extension",
				inputFull:     "/home/user1/projects/my-project/go-archetype.go",
				inputRelative: "my-project/go-archetype.go",
				expected:      "/home/user1/projects/my-project/project_name.go",
			},
			{
				name:          "project name is used several times",
				inputFull:     "/home/user1/projects/my-project/cmd/go-archetype/web/go-archetype.go",
				inputRelative: "my-project/cmd/go-archetype/web/go-archetype.go",
				expected:      "/home/user1/projects/my-project/cmd/project_name/web/project_name.go",
			},
			{
				name:          "project name is used outside of relative path",
				inputFull:     "/home/user1/projects/go-archetype/cmd/files/web/file.go",
				inputRelative: "cmd/files/web/file.go",
				expected:      "/home/user1/projects/go-archetype/cmd/files/web/file.go",
			},
			{
				name:          "project name is used inside and outside of relative path",
				inputFull:     "/home/user1/projects/go-archetype/cmd/files/web/go-archetype.go",
				inputRelative: "cmd/files/web/go-archetype.go",
				expected:      "/home/user1/projects/go-archetype/cmd/files/web/project_name.go",
			},
			{
				name:          "make sure, relative path is cleaned",
				inputFull:     "/home/user1/projects/go-archetype/cmd/files/web/go-archetype.go",
				inputRelative: ".//cmd/files/web/go-archetype.go",
				expected:      "/home/user1/projects/go-archetype/cmd/files/web/project_name.go",
			},
		}

		renamer := fileRenamer{
			name:        "project name",
			pattern:     "go-archetype",
			replacement: "project_name",
			files:       nil,
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				t.Parallel()

				input := types.File{
					Contents:     "",
					FullPath:     testCase.inputFull,
					RelativePath: testCase.inputRelative,
					Discarded:    false,
				}

				actual := renamer.Transform(input)
				assert.Equal(t, testCase.expected, actual.FullPath)
			})
		}
	})
}

func TestFileRenamer_Template(t *testing.T) {
	t.Parallel()

	t.Run("Template with empty vars", func(t *testing.T) {
		t.Parallel()

		testCases := []struct {
			name     string
			input    string
			expected string
		}{
			{
				name:     "empty replacement",
				input:    "",
				expected: "",
			},
			{
				name:     "filled replacement",
				input:    "awesome text here",
				expected: "awesome text here",
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				t.Parallel()

				renamer := &fileRenamer{replacement: testCase.input}

				err := renamer.Template(nil)
				require.NoError(t, err)
				assert.Equal(t, testCase.expected, renamer.replacement)
			})
		}
	})

	t.Run("Template with VARs", func(t *testing.T) {
		t.Parallel()

		testCases := []struct {
			name     string
			input    string
			expected string
		}{
			{
				name:     "empty replacement",
				input:    "",
				expected: "",
			},
			{
				name:     "filled replacement",
				input:    "awesome text here",
				expected: "awesome text here",
			},
			{
				name:     "template is used",
				input:    "{{ .FOO }}",
				expected: "BAR",
			},
			{
				name:     "2 templates are used",
				input:    "{{ .FOO }} - {{ .source }}",
				expected: "BAR - /home/user1",
			},
		}

		vars := map[string]string{
			"FOO":    "BAR",
			"source": "/home/user1",
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				t.Parallel()

				renamer := &fileRenamer{replacement: testCase.input}

				err := renamer.Template(vars)
				require.NoError(t, err)
				assert.Equal(t, testCase.expected, renamer.replacement)
			})
		}
	})

	t.Run("Handle errors", func(t *testing.T) {
		t.Parallel()

		testCases := []struct {
			name             string
			inputReplacement string
			inputVars        map[string]string
		}{
			{
				name:             "var is not defined",
				inputReplacement: "{{ .FOO }}",
				inputVars: map[string]string{
					"BAR": "FOO",
				},
			},
			{
				name:             "illegal format",
				inputReplacement: "{{ .FOO {{ }}",
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				t.Parallel()

				renamer := &fileRenamer{replacement: testCase.inputReplacement}

				err := renamer.Template(testCase.inputVars)
				assert.Error(t, err)
			})
		}
	})
}
