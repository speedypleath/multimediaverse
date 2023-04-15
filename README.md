# Multimediaverse

This is a collection of tools for creating and managing a multimedia files.

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
gocov test | gocov report
```

### Coverage

```bash
gocov test | gocov-xml > coverage/cov.xml
gocov test | gocov-html > coverage/report.html
```
