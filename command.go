package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func rootCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "text-template-cli TEMPLATE_FILE [INPUT_FILE...]",
		Short: "A simple CLI tool to render text templates",
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
			return renderTemplate(args[0], input)
		},
	}
}
