# hexlet-path-size

CLI utility for computing the size of a file or directory.

## Demo

[![asciicast](https://asciinema.org/a/Vo0Iixk524cQ80yw.svg)](https://asciinema.org/a/Vo0Iixk524cQ80yw)

## Installation

```bash
make build
```

## Usage

```bash
./bin/hexlet-path-size <path> [options]
```

**Flags:**

- `-r`, `--recursive` — calculate sizes recursively
- `-H`, `--human` — human readable format (KB, MB, GB)
- `-a`, `--all` — include hidden files

## Examples

```bash
./bin/hexlet-path-size README.md
./bin/hexlet-path-size -r -H .
./bin/hexlet-path-size -r -H -a /some/path
```

## Development

```bash
make test    # run tests
make lint    # run linter
make clean   # remove build artifacts
```
