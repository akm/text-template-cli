package main

import (
	"os"
	"path/filepath"
)

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
