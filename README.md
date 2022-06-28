# json2go

This CLI tool allows for the simple conversion of a JSON payload into ready-to-use Go types. The JSON payload can be supplied via file, STDIN, or URL fetching. The generation is able to correctly parse all top-level data types as well as deeply nested arrays and objects. Optional CLI arguments allow the customization of the outputted Go files.

### Installation

```bash
go install github.com/alehechka/json2go/cmd/json2go@latest
```

### Usage

From local JSON file:

```bash
json2go generate --file=example.json
```

From STDIN:

```bash
cat example.json | json2go generate
```

From URL fetching:

```bash
json2go generate --url="https://gorest.co.in/public/v2/users"
```

### CLI Arguments

| Argument         | Example                      | Type     | Purpose                                                                                                                                                                                            | Default    |
| ---------------- | ---------------------------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| File name        | `--file=example.json`        | `string` | Valid path to local file on disk                                                                                                                                                                   |
| URL              | `--url="https://example.com` | `string` | Valid URL that JSON can be fetched via GET request from.                                                                                                                                           |
| Root Object Name | `--root=RootObject`          | `string` | Name for top-level object in JSON payload                                                                                                                                                          | `Root`     |
| Package Name     | `--package=api`              | `string` | Name of package to generate types into. A nested package path is valid                                                                                                                             | `main`     |
| Output File Name | `--output`                   | `string` | The name of the file that is generated. If a file is provided as input, will use matching name unless explicitly provided. The ".go" extension is not required and will be automatically appended. | `types.go` |
| Time Format      | `--time=2006-01-02`          | `string` | Time format to use while parsing strings for potential time.Time variables. View time.Time constants for possible defaults: https://pkg.go.dev/time#pkg-constants                                  | `RFC3339`  |
| Debug logging    | `--debug`                    | `bool`   | Will output debugging console logs.                                                                                                                                                                | `false`    |
| Quiet            | `--quiet`                    | `bool`   | Will quiet fatal errors.                                                                                                                                                                           | `false`    |
| STDOUT           | `--out`                      | `bool`   | Instead of generating a Go file, will instead print the contents to STDOUT                                                                                                                         | `false`    |

### Local Development

While developing locally, the CLI can be used directly with the `go run` option as follows:

```bash
go run cmd/json2go/main.go generate --file=example.json <...args>
```
