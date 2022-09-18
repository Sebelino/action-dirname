# action-dirname
Given a set of file paths, outputs the set of directories containing those files.

## Inputs

### `files`

**Required** A JSON array of file paths.

## Outputs

### `directories`

A JSON array of directories containing the files.

## Example usage

```yaml
jobs:
  sample:
    runs-on: ubuntu-latest
    steps:
      - uses: sebelino/action-dirname@v1.0.5
        id: dirname
        with:
          files: '["main.tf", "versions.tf", "modules/vpc/sg.tf"]'
      - run: |
          # Will output: [".","modules/vpc"]
          echo ${{ toJson(steps.dirname.outputs.directories) }}
```

## Development

```bash
$ go run ./app.go '["main.tf", "versions.tf", "modules/vpc/sg.tf"]'
[".","modules/vpc"]
```
