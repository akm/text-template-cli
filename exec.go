package main

import (
	"bytes"
	"log"
	"os/exec"
	"text/template"
)

func execFuncs() template.FuncMap {
	return template.FuncMap{
		"stdout":      commandStdout,
		"stderr":      commandStderr,
		"combinedout": commandCombinedout,
		"shell":       shellExec,
	}
}

func commandStdout(command string, args ...string) string {
	cmd := exec.Command(command, args...)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func commandStderr(command string, args ...string) string {
	cmd := exec.Command(command, args...)
	stderr := bytes.Buffer{}
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	return stderr.String()
}

func commandCombinedout(command string, args ...string) string {
	cmd := exec.Command(command, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func shellExec(command string) string {
	cmd := exec.Command("sh", "-c", command)
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("sh -c %s failed with %+v\n", command, err)
	}
	return string(out)
}
