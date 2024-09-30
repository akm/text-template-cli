package main

import (
	"io"
	"os"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

func renderToStdout(templateFile string, input InputMap) error {
	return renderWithWriter(os.Stdout, templateFile, input)
}

func renderWithWriter(w io.Writer, templateFile string, input InputMap) error {
	templateData, err := os.ReadFile(templateFile)
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("template").Funcs(execFuncs()).Funcs(sprig.FuncMap()).Parse(string(templateData))
	if err != nil {
		panic(err)
	}

	if err := tmpl.Execute(w, input); err != nil {
		panic(err)
	}
	return nil
}
