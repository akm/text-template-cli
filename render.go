package main

import (
	"os"
	"text/template"
)

func renderTemplate(templateFile string, input InputMap) error {
	templateData, err := os.ReadFile(templateFile)
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("template").Parse(string(templateData))
	if err != nil {
		panic(err)
	}

	if err := tmpl.Execute(os.Stdout, input); err != nil {
		panic(err)
	}
	return nil
}
