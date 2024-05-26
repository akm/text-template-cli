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

You can use this application without any INPUT_FILE, but you can also pass JSON, YAML and/or .env files as INPUT FILE to render the template.

Usage:
  text-template-cli TEMPLATE_FILE [INPUT_FILE...] [flags]

Flags:
  -h, --help   help for text-template-cli
```

## Examples

See [examples](./examples/) for more details.

## If it doesnt't work

Check the path to sub-directory `bin` of `go env GOPATH` and ensure whether it is set in environment variable `PATH` .
If you use asdf, you might need to do `asdf reshim` .
