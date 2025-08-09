package code

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getFixturePath(name string) string {
	return filepath.Join("testdata/fixture", name)
}

func setupTestFile(t *testing.T, size int64) string {
	t.Helper()

	dir := t.TempDir()
	path := filepath.Join(dir, "test.file")

	f, err := os.Create(path)
	require.NoError(t, err)
	defer f.Close()

	err = f.Truncate(size)
	require.NoError(t, err)

	return path
}

func TestGetPathSize_EmptyDir(t *testing.T) {
	path := getFixturePath("empty_dir")

	got, err := GetPathSize(path, false, false, false)
	require.NoError(t, err)
	assert.Equal(t, "0B", got)
}

func TestGetPathSize_EmptyFile(t *testing.T) {
	path := getFixturePath("sizes/empty.txt")

	got, err := GetPathSize(path, false, false, false)
	require.NoError(t, err)
	assert.Equal(t, "0B", got)
}

func TestGetPathSize_SmallFile(t *testing.T) {
	path := getFixturePath("sizes/small.txt")

	got, err := GetPathSize(path, false, false, false)
	require.NoError(t, err)
	assert.Equal(t, "100B", got)
}

func TestGetPathSize_MediumFile(t *testing.T) {
	path := getFixturePath("sizes/medium.txt")

	got, err := GetPathSize(path, false, true, false)
	require.NoError(t, err)
	assert.Equal(t, "1.5KB", got)
}

func TestGetPathSize_LargeFile(t *testing.T) {
	path := getFixturePath("sizes/large.txt")

	got, err := GetPathSize(path, false, true, false)
	require.NoError(t, err)
	assert.Equal(t, "2.0KB", got)
}

func TestGetPathSize_NestedStructure(t *testing.T) {
	path := getFixturePath("nested")

	// Без рекурсии - только файл в корне
	got, err := GetPathSize(path, false, false, false)
	require.NoError(t, err)
	assert.Equal(t, "9B", got) // "root file" + newline

	// С рекурсией - все файлы
	got, err = GetPathSize(path, true, false, false)
	require.NoError(t, err)
	assert.Equal(t, "29B", got) // Сумма всех файлов без символов новой строки
}

func TestGetPathSize_HiddenFiles(t *testing.T) {
	path := getFixturePath("special")

	// Без скрытых файлов
	got, err := GetPathSize(path, true, false, false)
	require.NoError(t, err)
	assert.Equal(t, "15B", got) // Только unicode_файл.txt

	// Со скрытыми файлами
	got, err = GetPathSize(path, true, false, true)
	require.NoError(t, err)
	assert.Equal(t, "41B", got) // Все файлы, включая скрытые, без символов новой строки
}

func TestGetPathSize_Symlinks(t *testing.T) {
	t.Run("file link", func(t *testing.T) {
		path := getFixturePath("symlinks/link.txt")

		got, err := GetPathSize(path, false, false, false)
		require.NoError(t, err)
		assert.Equal(t, "13B", got) // "original file" + newline
	})

	t.Run("directory link", func(t *testing.T) {
		path := getFixturePath("symlinks/dir_link")

		got, err := GetPathSize(path, true, false, false)
		require.NoError(t, err)
		assert.Equal(t, "29B", got) // То же, что и в тесте NestedStructure
	})
}

func TestGetPathSize_NonExistentPath(t *testing.T) {
	path := getFixturePath("non_existent")

	_, err := GetPathSize(path, false, false, false)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error processing")
}

func TestGetPathSize_UnicodeFileName(t *testing.T) {
	path := getFixturePath("special/unicode_файл.txt")

	got, err := GetPathSize(path, false, false, false)
	require.NoError(t, err)
	assert.Equal(t, "15B", got) // "unicode content" + newline
}

func TestGetPathSize_HumanReadableSizes(t *testing.T) {
	tests := []struct {
		name     string
		size     int64
		expected string
	}{
		{"zero bytes", 0, "0B"},
		{"1 byte", 1, "1B"},
		{"1023 bytes", 1023, "1023B"},
		{"1024 bytes", 1024, "1.0KB"},
		{"1536 bytes", 1536, "1.5KB"},
		{"1MB", 1024 * 1024, "1.0MB"},
		{"1.5MB", 1024 * 1024 * 3 / 2, "1.5MB"},
		{"1GB", 1024 * 1024 * 1024, "1.0GB"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			path := setupTestFile(t, tc.size)
			got, err := GetPathSize(path, false, true, false)
			require.NoError(t, err)
			assert.Equal(t, tc.expected, got)
		})
	}
}
