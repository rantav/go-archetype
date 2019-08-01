package transformer

import (
	"bytes"
	tt "text/template"

	"github.com/pkg/errors"
)

func template(text string, vars map[string]string) (string, error) {
	tmpl, err := tt.New("t").Parse(text)
	if err != nil {
		return "", errors.Wrap(err, "error creating the text template")
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, vars)
	if err != nil {
		return "", errors.Wrap(err, "error templating")
	}
	return buf.String(), nil
}
