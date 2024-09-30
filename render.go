package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

func renderToStdout(templateFile string, input InputMap) error {
	return renderWithWriter(os.Stdout, templateFile, input)
}

func renderToFile(templateFile string, input InputMap, dest string) error {
	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer f.Close()
	return renderWithWriter(f, templateFile, input)
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

func renderOrCopy(srcFile, destPath string, templateExts []string, input InputMap) error {
	isTemplate := false
	for _, ext := range templateExts {
		if strings.HasSuffix(destPath, ext) {
			isTemplate = true
			destPath = strings.TrimSuffix(destPath, ext)
			break
		}
	}

	// Write the rendered content to the destination directory
	if err := os.MkdirAll(filepath.Dir(destPath), 0o755); err != nil {
		return err
	}

	if isTemplate {
		return renderToFile(srcFile, input, destPath)
	} else {
		// Copy non-template files directly
		return copyFile(srcFile, destPath)
	}
}

// copyFile copies a file from src to dst
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	return err
}
