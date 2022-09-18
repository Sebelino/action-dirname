# action-dirname
Given a JSON array of file paths, outputs a JSON array of directories of those files.

## Inputs

### `files`

**Required** A JSON array of file paths.

## Outputs

### `directories`

A JSON array of directories containing the files.

## Example usage

```bash
$ go run ./app.go '["main.tf", "versions.tf", "modules/vpc/sg.tf"]'
[".","modules/vpc"]
```
