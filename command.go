package main

import (
	"fmt"
	"os"
	"path/filepath"

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
