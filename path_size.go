package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// NOTE: Экспортируется только одна функция
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := getSize(path, recursive, all)
	if err != nil {
		return "", fmt.Errorf("error processing %s: %v", path, err)
	}
	result := formatSize(size, human)
	return result, nil
}

// getSize returns the total size in bytes of the file or directory at path.
// If recursive is true and path is a directory, it sums sizes recursively.
// Otherwise, it sums sizes of files directly under the directory.
// Hidden entries (names starting with ".") are included only when all is true.
func getSize(path string, recursive, all bool) (int64, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}
	if !info.IsDir() {
		return info.Size(), nil
	}
	var total int64
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}
	for _, entry := range entries {
		if !all && strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		entryPath := filepath.Join(path, entry.Name())
		entryInfo, err := entry.Info()
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: can't stat %s: %v\n", entryPath, err)
			continue
		}
		if entry.IsDir() {
			if recursive {
				size, err := getSize(entryPath, recursive, all)
				if err != nil {
					continue
				}
				total += size
			}
		} else {
			total += entryInfo.Size()
		}
	}
	return total, nil
}

func formatSize(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", size)
	}

	const (
		KB = 1024
		MB = 1024 * KB
		GB = 1024 * MB
	)

	switch {
	case size < KB:
		return fmt.Sprintf("%dB", size)
	case size < MB:
		return fmt.Sprintf("%.1fKB", float64(size)/KB)
	case size < GB:
		return fmt.Sprintf("%.1fMB", float64(size)/MB)
	default:
		return fmt.Sprintf("%.1fGB", float64(size)/GB)
	}
}
