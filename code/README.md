## Overview

hexlet-path-size is a CLI utility to calculate the size of files or directories. It supports recursive size calculation, human-readable formats, and inclusion of hidden files.

## Features

- **Recursive Size Calculation**: Calculate the size of all nested files and directories.
- **Human-Readable Format**: Display sizes in KB, MB, GB, etc.
- **Include Hidden Files**: Optionally include hidden files and directories.

## Usage

### Basic Usage

```bash
hexlet-path-size <path>
```

### Options

- `--recursive`, `-r`: Calculate size recursively for directories.
- `--human`, `-H`: Display size in human-readable format.
- `--all`, `-a`: Include hidden files and directories.

### Examples

```bash
hexlet-path-size /path/to/directory
hexlet-path-size /path/to/directory --recursive
hexlet-path-size /path/to/directory --human
hexlet-path-size /path/to/directory --recursive --all
```

## Installation

1. Clone the repository:
   ```bash
   git clone <repository_url>
   ```
2. Navigate to the project directory:
   ```bash
   cd hexlet-path-size
   ```
3. Build the project:
   ```bash
   go build -o bin/hexlet-path-size main.go
   ```

## Testing

Run tests using:

```bash
make test
```
