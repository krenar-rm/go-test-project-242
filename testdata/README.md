# Test Data Structure

This directory contains test fixtures for the path-size utility tests.

## Directory Structure

```
testdata/
├── README.md
└── fixture/
    ├── empty_dir/          # Empty directory for testing
    ├── sizes/              # Files of different sizes
    │   ├── empty.txt      # 0 bytes
    │   ├── small.txt      # 100 bytes
    │   ├── medium.txt     # 1.5 KB
    │   └── large.txt      # 2.0 KB
    ├── nested/            # Directory with nested structure
    │   ├── level1/
    │   │   ├── level2/
    │   │   │   └── deep.txt
    │   │   └── file.txt
    │   └── file.txt
    ├── special/           # Special cases
    │   ├── .hidden       # Hidden file
    │   ├── .hidden_dir/  # Hidden directory
    │   │   └── file.txt
    │   └── unicode_файл.txt # File with Unicode name
    └── symlinks/          # Symbolic links
        ├── file.txt      # Original file
        ├── link.txt      # Symlink to file.txt
        └── dir_link      # Symlink to nested directory
```

## File Contents

- All text files contain predictable content (repeated characters or numbers)
- Files in nested directories contain their path as content
- Size-specific files contain exact number of bytes as specified in their names