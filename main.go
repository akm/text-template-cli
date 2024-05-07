package main

import (
	"os"
	"strings"
	"text/template"
)

func main() {
	if len(os.Args) < 2 {
		panic("no template file specified")
	}

	templateData, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("template").Parse(string(templateData))
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, getEnvs())
	if err != nil {
		panic(err)
	}
}

func getEnvs() map[string]string {
	envs := make(map[string]string)
	for _, env := range os.Environ() {
		pair := strings.Split(env, "=")
		envs[pair[0]] = pair[1]
	}
	return envs
}
