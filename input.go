package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type InputMap map[string]interface{}

func buildInput(inputFiles []string) (InputMap, error) {
	result := InputMap{}
	for k, v := range getEnvs() {
		result[k] = v
	}
	for _, inputFile := range inputFiles {
		if err := loadFile(result, inputFile); err != nil {
			return nil, err
		}
	}
	return result, nil
}

func getEnvs() map[string]string {
	envs := make(map[string]string)
	for _, env := range os.Environ() {
		pair := strings.Split(env, "=")
		envs[pair[0]] = pair[1]
	}
	return envs
}

func loadFile(dest InputMap, fpath string) error {
	switch filepath.Ext(fpath) {
	case ".json":
		return loadJSONFile(dest, fpath)
	case ".env":
		return loadEnvFile(dest, fpath)
	default:
		return fmt.Errorf("unsupported file type: %s", fpath)
	}
}

func loadJSONFile(dest InputMap, fpath string) error {
	f, err := os.Open(fpath)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	var data InputMap
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	for k, v := range data {
		oldVal, ok := dest[k]
		if ok {
			fmt.Fprintf(os.Stderr, "Overwriting %s: %s -> %s by %s\n", k, oldVal, v, fpath)
		}
		dest[k] = v
	}

	return nil
}

func loadEnvFile(dest InputMap, fpath string) error {
	f, err := os.Open(fpath)
	if err != nil {
		return err
	}
	defer f.Close()

	envs := make(InputMap)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			envs[parts[0]] = ""
		} else {
			envs[parts[0]] = parts[1]
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	for k, v := range envs {
		oldVal, ok := dest[k]
		if ok {
			fmt.Printf("[WARN] Overwriting %s: %s -> %s by %s\n", k, oldVal, v, fpath)
		}
		dest[k] = v
	}

	return nil
}
