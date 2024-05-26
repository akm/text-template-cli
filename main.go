package main

import (
	"fmt"
	"os"
)

func main() {
	rootCmd := rootCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}
