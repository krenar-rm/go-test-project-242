package pathsize

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Options struct {
	Recursive bool
	Human     bool
	All       bool
}

func GetPathSize(path string, opts Options) (string, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return "", fmt.Errorf("stat %s: %w", path, err)
	}

	bytes, err := computeSize(path, info, opts)
	if err != nil {
		return "", err
	}

	return formatBytes(bytes, opts.Human), nil
}

func computeSize(path string, info fs.FileInfo, opts Options) (int64, error) {
	if !info.IsDir() {
		return info.Size(), nil
	}
	if opts.Recursive {
		return walkDirSize(path, opts.All)
	}
	return shallowDirSize(path, opts.All)
}

func walkDirSize(root string, includeHidden bool) (int64, error) {
	var total int64
	err := filepath.WalkDir(root, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: %v\n", err)
			return nil
		}
		if p == root {
			return nil
		}
		if !includeHidden && strings.HasPrefix(d.Name(), ".") {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if d.IsDir() {
			return nil
		}
		info, err := d.Info()
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: %v\n", err)
			return nil
		}
		total += info.Size()
		return nil
	})
	return total, err
}

func shallowDirSize(root string, includeHidden bool) (int64, error) {
	entries, err := os.ReadDir(root)
	if err != nil {
		return 0, fmt.Errorf("read dir %s: %w", root, err)
	}

	var total int64
	for _, e := range entries {
		if !includeHidden && strings.HasPrefix(e.Name(), ".") {
			continue
		}
		if e.IsDir() {
			continue
		}
		info, err := e.Info()
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: %v\n", err)
			continue
		}
		total += info.Size()
	}
	return total, nil
}

var sizeUnits = []struct {
	threshold int64
	suffix    string
}{
	{1 << 30, "GB"},
	{1 << 20, "MB"},
	{1 << 10, "KB"},
}

func formatBytes(n int64, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", n)
	}
	for _, u := range sizeUnits {
		if n >= u.threshold {
			return fmt.Sprintf("%.1f%s", float64(n)/float64(u.threshold), u.suffix)
		}
	}
	return fmt.Sprintf("%dB", n)
}
