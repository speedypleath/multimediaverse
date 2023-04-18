# Multimediaverse

This is a collection of tools for creating and managing a multimedia files.

## Installation

```bash
go install
```

## Usage

```bash
Usage:
  multimediaverse [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  env         Configure envirmonment variables
  help        Help about any command
  video       Video editing tools

Flags:
      --config string   config file (default is $HOME/.multimediaverse.yaml)
  -h, --help            help for multimediaverse
  -t, --toggle          Help message for toggle

Use "multimediaverse [command] --help" for more information about a command.
```

## Development

### Requirements

* [gocov](https://github.com/axw/gocov) (for coverage)
* [gocov-xml](https://github.com/AlekSi/gocov-xml) (for IDE integration)
* [gocov-html](https://github.com/matm/gocov-html) (for coverage visualization)

### Build

```bash
go build -o build/multimediaverse multimediaverse
```

### Run

```bash
go run multimediaverse
```

### Test

```bash
gocov test ./...
```

### Coverage

```bash
gocov test ./... | gocov-xml > coverage/cov.xml
gocov test ./... | gocov-html > coverage/report.html
```
