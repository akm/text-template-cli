package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

func rootCommand() *cobra.Command {
	var destDirectory string
	var templateExts []string
	r := &cobra.Command{
		Version: Version,
		Use:     "text-template-cli TEMPLATE_FILE [INPUT_FILE...]",
		Short:   "A simple CLI tool to render text templates",
		Long: `A simple CLI tool to render text templates.
This application is a tool to render text templates using Go's text/template package.
See https://pkg.go.dev/text/template for more information about template file.

You can use this application without any INPUT_FILE, but you can also pass JSON, YAML and/or .env files as INPUT FILE to render the template.`,
		Args: func(_ *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("TEMPLATE_FILE is required")
			}
			return nil
		},
		RunE: func(_ *cobra.Command, args []string) error {
			input, err := buildInput(args[1:])
			if err != nil {
				return err
			}

			f, err := os.Open(args[0])
			if err != nil {
				return err
			}
			fileInfo, err := f.Stat()
			if err != nil {
				return err
			}
			if fileInfo.IsDir() {
				if destDirectory == "" {
					return fmt.Errorf("output-directory is required when TEMPLATE_FILE is a directory")
				}
				return renderDirectory(args[0], destDirectory, templateExts, input)
			} else {
				return renderTemplate(args[0], input)
			}
		},
	}
	r.Flags().StringVarP(&destDirectory, "output-directory", "o", "", "Output directory")
	r.Flags().StringSliceVarP(&templateExts, "template-ext", "e", []string{".tmpl"}, "Template file extensions")
	return r
}

func renderDirectory(srcDir, destDir string, templateExts []string, input InputMap) error {
	var files []string

	// Walk through the directory and collect file paths
	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	// Process each file
	for _, file := range files {
		relPath, err := filepath.Rel(srcDir, file)
		if err != nil {
			return err
		}
		destPath := filepath.Join(destDir, relPath)

		if isTemplateFile(file, templateExts) {
			// Parse and execute the template
			tmpl, err := template.ParseFiles(file)
			if err != nil {
				return err
			}

			var renderedTemplate bytes.Buffer
			err = tmpl.Execute(&renderedTemplate, input)
			if err != nil {
				return err
			}

			// Remove template extension from the destination file name
			for _, ext := range templateExts {
				if strings.HasSuffix(destPath, ext) {
					destPath = strings.TrimSuffix(destPath, ext)
					break
				}
			}

			// Write the rendered content to the destination directory
			err = os.MkdirAll(filepath.Dir(destPath), os.ModePerm)
			if err != nil {
				return err
			}
			f, err := os.Create(destPath)
			if err != nil {
				return err
			}
			defer f.Close()

			if _, err = f.Write(renderedTemplate.Bytes()); err != nil {
				return err
			}
		} else {
			// Copy non-template files directly
			err = copyFile(file, destPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// isTemplateFile checks if a file is a template file based on its extension
func isTemplateFile(filename string, templateExts []string) bool {
	for _, ext := range templateExts {
		if strings.HasSuffix(filename, ext) {
			return true
		}
	}
	return false
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
