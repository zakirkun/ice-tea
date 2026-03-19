package scanner

import (
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
)

// Walker discovers source code files while respecting exclusion rules
type Walker struct {
	ExcludeDirs       map[string]bool
	ExcludeFiles      []string // glob patterns
	ExcludeExtensions map[string]bool
	Languages         []string // optional language filter
	logger            *zap.SugaredLogger
}

// WalkerConfig configures the file walker
type WalkerConfig struct {
	ExcludeDirs       []string
	ExcludeFiles      []string
	ExcludeExtensions []string
	Languages         []string
}

// NewWalker creates a new file walker with the given config
func NewWalker(cfg WalkerConfig, logger *zap.SugaredLogger) *Walker {
	excludeDirs := make(map[string]bool)
	for _, d := range cfg.ExcludeDirs {
		excludeDirs[d] = true
	}

	excludeExts := make(map[string]bool)
	for _, e := range cfg.ExcludeExtensions {
		if !strings.HasPrefix(e, ".") {
			e = "." + e
		}
		excludeExts[strings.ToLower(e)] = true
	}

	return &Walker{
		ExcludeDirs:       excludeDirs,
		ExcludeFiles:      cfg.ExcludeFiles,
		ExcludeExtensions: excludeExts,
		Languages:         cfg.Languages,
		logger:            logger,
	}
}

// DiscoveredFile represents a file found by the walker
type DiscoveredFile struct {
	Path     string // absolute path
	RelPath  string // path relative to scan root
	Language string // detected language
	Size     int64  // file size in bytes
}

// Walk discovers all scannable files under the given root path
func (w *Walker) Walk(root string) ([]DiscoveredFile, error) {
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return nil, err
	}

	info, err := os.Stat(absRoot)
	if err != nil {
		return nil, err
	}

	// If root is a single file, check it directly
	if !info.IsDir() {
		lang := DetectLanguage(absRoot)
		if lang == "" {
			return nil, nil
		}
		return []DiscoveredFile{{
			Path:     absRoot,
			RelPath:  filepath.Base(absRoot),
			Language: lang,
			Size:     info.Size(),
		}}, nil
	}

	var files []DiscoveredFile

	err = filepath.WalkDir(absRoot, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			w.logger.Warnw("Error accessing path", "path", path, "error", err)
			return nil // skip errors, continue walking
		}

		relPath, _ := filepath.Rel(absRoot, path)

		// Skip excluded directories
		if d.IsDir() {
			dirName := d.Name()
			if w.ExcludeDirs[dirName] {
				w.logger.Debugw("Skipping excluded directory", "dir", relPath)
				return filepath.SkipDir
			}
			// Skip hidden directories (starting with .)
			if strings.HasPrefix(dirName, ".") && dirName != "." {
				return filepath.SkipDir
			}
			return nil
		}

		// Skip excluded extensions
		ext := strings.ToLower(filepath.Ext(path))
		if w.ExcludeExtensions[ext] {
			return nil
		}

		// Skip excluded file patterns
		fileName := d.Name()
		for _, pattern := range w.ExcludeFiles {
			matched, _ := filepath.Match(pattern, fileName)
			if matched {
				w.logger.Debugw("Skipping excluded file", "file", relPath, "pattern", pattern)
				return nil
			}
		}

		// Detect language
		lang := DetectLanguage(path)
		if lang == "" {
			return nil // unsupported file type
		}

		// Apply language filter
		if len(w.Languages) > 0 && !w.isLanguageIncluded(lang) {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return nil
		}

		// Skip empty files
		if info.Size() == 0 {
			return nil
		}

		files = append(files, DiscoveredFile{
			Path:     path,
			RelPath:  relPath,
			Language: lang,
			Size:     info.Size(),
		})

		return nil
	})

	return files, err
}

func (w *Walker) isLanguageIncluded(lang string) bool {
	for _, l := range w.Languages {
		if strings.EqualFold(l, lang) {
			return true
		}
	}
	return false
}

// DetectLanguage returns the language based on file extension
func DetectLanguage(path string) string {
	ext := strings.ToLower(filepath.Ext(path))
	lang, ok := extensionToLanguage[ext]
	if !ok {
		return ""
	}
	return lang
}

// extensionToLanguage maps file extensions to language identifiers
var extensionToLanguage = map[string]string{
	// Go
	".go": "go",

	// JavaScript / TypeScript
	".js":  "javascript",
	".jsx": "javascript",
	".ts":  "typescript",
	".tsx": "typescript",
	".mjs": "javascript",
	".cjs": "javascript",

	// Python
	".py":  "python",
	".pyw": "python",

	// Java
	".java": "java",

	// PHP
	".php":  "php",
	".php3": "php",
	".php4": "php",
	".php5": "php",

	// Ruby
	".rb":  "ruby",
	".erb": "ruby",

	// Rust
	".rs": "rust",

	// C / C++
	".c":   "c",
	".h":   "c",
	".cpp": "cpp",
	".cc":  "cpp",
	".cxx": "cpp",
	".hpp": "cpp",
	".hxx": "cpp",

	// Shell
	".sh":   "shell",
	".bash": "shell",
	".zsh":  "shell",

	// YAML (for config scanning)
	".yaml": "yaml",
	".yml":  "yaml",

	// Dart (Flutter)
	".dart": "dart",

	// Kotlin
	".kt":  "kotlin",
	".kts": "kotlin",

	// Zig
	".zig": "zig",

	// Perl
	".pl": "perl",
	".pm": "perl",

	// Elixir
	".ex":  "elixir",
	".exs": "elixir",

	// Dockerfile
	// Note: Dockerfile has no extension, handled separately
}

// SupportedLanguages returns a list of all supported language identifiers
func SupportedLanguages() []string {
	seen := make(map[string]bool)
	var langs []string
	for _, lang := range extensionToLanguage {
		if !seen[lang] {
			seen[lang] = true
			langs = append(langs, lang)
		}
	}
	return langs
}
