package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Единицы измерения для размеров файлов
const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

// Options содержит параметры для вычисления размера
type Options struct {
	Recursive bool // рекурсивный подсчет для директорий
	Human     bool // человекочитаемый формат вывода
	All       bool // включать скрытые файлы
}

// GetPathSize возвращает размер файла или директории в виде строки.
// Если path указывает на файл, возвращается его размер.
// Если path указывает на директорию:
//   - при recursive=true подсчитывается размер всех файлов в директории и поддиректориях
//   - при recursive=false подсчитывается размер только файлов в текущей директории
//   - при all=true учитываются скрытые файлы (начинающиеся с точки)
//   - при human=true размер форматируется в человекочитаемом виде (B, KB, MB, GB)
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	opts := Options{
		Recursive: recursive,
		Human:     human,
		All:       all,
	}

	size, err := getSize(path, opts)
	if err != nil {
		return "", fmt.Errorf("failed to get size of %q: %w", path, err)
	}

	return formatSize(size, opts.Human), nil
}

// getSize возвращает размер в байтах для указанного пути.
func getSize(path string, opts Options) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, fmt.Errorf("failed to stat %q: %w", path, err)
	}

	if !info.IsDir() {
		return info.Size(), nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, fmt.Errorf("failed to read directory %q: %w", path, err)
	}

	var total int64
	for _, entry := range entries {
		if !opts.All && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		entryPath := filepath.Join(path, entry.Name())
		info, err := entry.Info()
		if err != nil {
			// Логируем ошибку, но продолжаем обработку
			fmt.Fprintf(os.Stderr, "Warning: failed to get info for %q: %v\n", entryPath, err)
			continue
		}

		if info.IsDir() {
			if opts.Recursive {
				size, err := getSize(entryPath, opts)
				if err != nil {
					// Логируем ошибку, но продолжаем обработку
					fmt.Fprintf(os.Stderr, "Warning: failed to get size of %q: %v\n", entryPath, err)
					continue
				}
				total += size
			}
		} else {
			total += info.Size()
		}
	}

	return total, nil
}

// formatSize преобразует размер в байтах в человекочитаемый формат
func formatSize(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", size)
	}

	switch {
	case size >= GB:
		return fmt.Sprintf("%.1fGB", float64(size)/float64(GB))
	case size >= MB:
		return fmt.Sprintf("%.1fMB", float64(size)/float64(MB))
	case size >= KB:
		return fmt.Sprintf("%.1fKB", float64(size)/float64(KB))
	default:
		return fmt.Sprintf("%dB", size)
	}
}
