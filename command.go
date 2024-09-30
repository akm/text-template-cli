package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func rootCommand() *cobra.Command {
	var destDirectory string
	var templateExts []string
	r := &cobra.Command{
		Version: Version,
		Use:     "text-template-cli SOURCE [INPUT_FILE...]",
		Short:   "A simple CLI tool to render text templates",
		Long: `A simple CLI tool to render text templates.
This application is a tool to render text templates using Go's text/template package.
See https://pkg.go.dev/text/template for more information about template file.

SOURCE must be a file or a directory.
If SOURCE is a file, the rendered content will be written to the standard output, 
or will be written a file in the output directory when --output-directory is given.
If SOURCE is a directory, the rendered content will be written to the output directory.

You can use this application without any INPUT_FILE, but you can also pass JSON, YAML and/or .env files as INPUT FILE to render the template.`,
		Args: func(_ *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("SOURCE is required")
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
					return fmt.Errorf("output-directory is required when SOURCE is a directory")
				}
				return processDirectory(args[0], destDirectory, templateExts, input)
			} else if destDirectory != "" {
				return renderOrCopy(args[0], filepath.Join(destDirectory, filepath.Base(args[0])), templateExts, input)
			} else {
				return renderToStdout(args[0], input)
			}
		},
	}
	r.Flags().StringVarP(&destDirectory, "output-directory", "o", "", "Output directory")
	r.Flags().StringSliceVarP(&templateExts, "template-ext", "e", []string{".tmpl"}, "Template file extensions")
	return r
}

func processDirectory(srcDir, destDir string, templateExts []string, input InputMap) error {
	var srcFiles []string

	// Walk through the directory and collect file paths
	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			srcFiles = append(srcFiles, path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	// Process each file
	for _, srcFile := range srcFiles {
		srcRelPath, err := filepath.Rel(srcDir, srcFile)
		if err != nil {
			return err
		}
		destPath := filepath.Join(destDir, srcRelPath)

		if err := renderOrCopy(srcFile, destPath, templateExts, input); err != nil {
			return err
		}
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
