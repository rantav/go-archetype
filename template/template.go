package template

import (
	"bytes"
	ht "html/template"
	tt "text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/pkg/errors"
)

func Execute(text string, vars map[string]string) (string, error) {
	tmpl, err := tt.New("t").Funcs(textmap(sprig.FuncMap())).Parse(text)
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

func textmap(h ht.FuncMap) tt.FuncMap {
	m := make(tt.FuncMap)
	for k, v := range h {
		m[k] = v
	}
	return m
}
