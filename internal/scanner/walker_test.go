package scanner

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func testLogger() *zap.SugaredLogger {
	logger, _ := zap.NewDevelopment()
	return logger.Sugar()
}

func TestDetectLanguage(t *testing.T) {
	tests := []struct {
		path string
		want string
	}{
		{"main.go", "go"},
		{"app.js", "javascript"},
		{"app.jsx", "javascript"},
		{"index.ts", "typescript"},
		{"index.tsx", "typescript"},
		{"script.py", "python"},
		{"Main.java", "java"},
		{"index.php", "php"},
		{"app.rb", "ruby"},
		{"main.rs", "rust"},
		{"main.c", "c"},
		{"main.cpp", "cpp"},
		{"deploy.sh", "shell"},
		{"config.yaml", "yaml"},
		{"README.md", ""},       // unsupported
		{"image.png", ""},       // unsupported
		{"data.json", ""},       // not yet supported
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			got := DetectLanguage(tt.path)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestWalkerExclusions(t *testing.T) {
	// Create temp directory structure
	tmpDir := t.TempDir()

	// Create test files
	createFile(t, tmpDir, "main.go", "package main")
	createFile(t, tmpDir, "utils.go", "package main")
	createFile(t, tmpDir, "app.js", "console.log('hello')")
	createFile(t, tmpDir, "README.md", "# Readme")
	createFile(t, tmpDir, "image.png", "fake png")

	// Create subdirectories
	createFile(t, filepath.Join(tmpDir, "src"), "handler.go", "package src")
	createFile(t, filepath.Join(tmpDir, "vendor"), "dep.go", "package dep")
	createFile(t, filepath.Join(tmpDir, "node_modules"), "lib.js", "module.exports = {}")
	createFile(t, filepath.Join(tmpDir, ".git"), "config", "git config")

	walker := NewWalker(WalkerConfig{
		ExcludeDirs:       []string{"vendor", "node_modules"},
		ExcludeFiles:      []string{"*.min.js"},
		ExcludeExtensions: []string{".md", ".png"},
	}, testLogger())

	files, err := walker.Walk(tmpDir)
	require.NoError(t, err)

	// Check results
	var paths []string
	for _, f := range files {
		paths = append(paths, f.RelPath)
	}

	assert.Contains(t, paths, "main.go")
	assert.Contains(t, paths, "utils.go")
	assert.Contains(t, paths, "app.js")
	assert.Contains(t, paths, filepath.Join("src", "handler.go"))

	// Excluded items should not be present
	assert.NotContains(t, paths, "README.md")              // excluded extension
	assert.NotContains(t, paths, "image.png")               // excluded extension
	assert.NotContains(t, paths, filepath.Join("vendor", "dep.go"))       // excluded dir
	assert.NotContains(t, paths, filepath.Join("node_modules", "lib.js")) // excluded dir
}

func TestWalkerLanguageFilter(t *testing.T) {
	tmpDir := t.TempDir()

	createFile(t, tmpDir, "main.go", "package main")
	createFile(t, tmpDir, "app.js", "console.log('hello')")
	createFile(t, tmpDir, "script.py", "print('hello')")

	walker := NewWalker(WalkerConfig{
		Languages: []string{"go"},
	}, testLogger())

	files, err := walker.Walk(tmpDir)
	require.NoError(t, err)

	assert.Len(t, files, 1)
	assert.Equal(t, "go", files[0].Language)
}

func TestWalkerSingleFile(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "main.go")
	createFile(t, tmpDir, "main.go", "package main")

	walker := NewWalker(WalkerConfig{}, testLogger())
	files, err := walker.Walk(filePath)
	require.NoError(t, err)

	assert.Len(t, files, 1)
	assert.Equal(t, "go", files[0].Language)
}

func TestWalkerEmptyFileSkipped(t *testing.T) {
	tmpDir := t.TempDir()
	createFile(t, tmpDir, "empty.go", "")

	walker := NewWalker(WalkerConfig{}, testLogger())
	files, err := walker.Walk(tmpDir)
	require.NoError(t, err)

	assert.Len(t, files, 0)
}

func TestSupportedLanguages(t *testing.T) {
	langs := SupportedLanguages()
	assert.True(t, len(langs) > 0)
	assert.Contains(t, langs, "go")
	assert.Contains(t, langs, "javascript")
	assert.Contains(t, langs, "python")
}

// helper: create a file in dir with given name and content
func createFile(t *testing.T, dir, name, content string) {
	t.Helper()
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
}
