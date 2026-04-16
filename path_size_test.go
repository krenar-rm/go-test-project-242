package code_test

import (
	"os"
	"path/filepath"
	"testing"

	"code"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const fixtureDir = "testdata/fixture"

func TestSingleFile(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		human    bool
		expected string
	}{
		{"empty", "sizes/empty.txt", false, "0B"},
		{"small", "sizes/small.txt", false, "100B"},
		{"medium bytes", "sizes/medium.txt", false, "1536B"},
		{"medium human", "sizes/medium.txt", true, "1.5KB"},
		{"large bytes", "sizes/large.txt", false, "2048B"},
		{"large human", "sizes/large.txt", true, "2.0KB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := code.GetPathSize(filepath.Join(fixtureDir, tt.path), false, tt.human, false)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestDirectoryNonRecursive(t *testing.T) {
	got, err := code.GetPathSize(filepath.Join(fixtureDir, "nested"), false, false, false)
	require.NoError(t, err)
	assert.Equal(t, "10B", got)
}

func TestDirectoryRecursive(t *testing.T) {
	got, err := code.GetPathSize(filepath.Join(fixtureDir, "nested"), true, false, false)
	require.NoError(t, err)
	assert.Equal(t, "30B", got)
}

func TestHiddenFiles(t *testing.T) {
	dir := filepath.Join(fixtureDir, "special")

	got, err := code.GetPathSize(dir, true, false, false)
	require.NoError(t, err)
	assert.Equal(t, "29B", got)

	got, err = code.GetPathSize(dir, true, false, true)
	require.NoError(t, err)
	assert.Equal(t, "43B", got)
}

func TestUnicodeFileName(t *testing.T) {
	got, err := code.GetPathSize(filepath.Join(fixtureDir, "special", "unicode_файл.txt"), false, false, false)
	require.NoError(t, err)
	assert.Equal(t, "13B", got)
}

func TestNonExistentPath(t *testing.T) {
	_, err := code.GetPathSize("/no/such/path", false, false, false)
	assert.Error(t, err)
}

func TestEmptyDir(t *testing.T) {
	got, err := code.GetPathSize(t.TempDir(), true, false, false)
	require.NoError(t, err)
	assert.Equal(t, "0B", got)
}

func TestSymlink(t *testing.T) {
	dir := t.TempDir()
	real := filepath.Join(dir, "real.txt")
	link := filepath.Join(dir, "link.txt")

	require.NoError(t, os.WriteFile(real, []byte("hello world"), 0644))
	require.NoError(t, os.Symlink(real, link))

	got, err := code.GetPathSize(link, false, false, false)
	require.NoError(t, err)
	assert.NotEmpty(t, got)
}

func TestHumanUnits(t *testing.T) {
	cases := []struct {
		size     int64
		expected string
	}{
		{0, "0B"},
		{1, "1B"},
		{1023, "1023B"},
		{1024, "1.0KB"},
		{1536, "1.5KB"},
		{1 << 20, "1.0MB"},
		{1<<20 + 1<<19, "1.5MB"},
		{1 << 30, "1.0GB"},
	}
	for _, c := range cases {
		dir := t.TempDir()
		p := filepath.Join(dir, "f")
		f, err := os.Create(p)
		require.NoError(t, err)
		if c.size > 0 {
			require.NoError(t, f.Truncate(c.size))
		}
		f.Close()

		got, err := code.GetPathSize(p, false, true, false)
		require.NoError(t, err)
		assert.Equal(t, c.expected, got)
	}
}
