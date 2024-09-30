# text-template-cli

CLI for [text/template](https://pkg.go.dev/text/template) of golang

## Install

```
go install github.com/akm/text-template-cli@latest
```

### Usage

```
$ text-template-cli --help
A simple CLI tool to render text templates.
This application is a tool to render text templates using Go's text/template package.
See https://pkg.go.dev/text/template for more information about template file.

SOURCE must be a file or a directory.
If SOURCE is a file, the rendered content will be written to the standard output, 
or will be written a file in the output directory when --output-directory is given.
If SOURCE is a directory, the rendered content will be written to the output directory.

You can use this application without any INPUT_FILE, but you can also pass JSON, YAML and/or .env files as INPUT FILE to render the template.

Usage:
  text-template-cli SOURCE [INPUT_FILE...] [flags]

Flags:
  -h, --help                      help for text-template-cli
  -o, --output-directory string   Output directory
  -e, --template-ext strings      Template file extensions (default [.tmpl])
  -v, --version                   version for text-template-cli
```

## Functions

- [Sprig](https://masterminds.github.io/sprig/) functions are available.
- External command functions are available.

### External command functions

| name        | Executeion                               | Output            |
| ----------- | ---------------------------------------- | ----------------- |
| stdout      | Execute given command                    | stdout            |
| stderr      | Execute given command                    | stderr            |
| combinedout | Execute given command                    | stdout and stderr |
| shell       | Execute `sh` with `-c` and given command | stdout            |

See [examples/README.md.tmpl](examples/README.md.tmpl).

## Examples

See [examples](./examples/) for more details.

## If it doesnt't work

Check the path to sub-directory `bin` of `go env GOPATH` and ensure whether it is set in environment variable `PATH` .
If you use asdf, you might need to do `asdf reshim` .
