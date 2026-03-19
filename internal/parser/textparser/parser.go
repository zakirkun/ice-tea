package textparser

import (
	"strings"

	p "github.com/zakirkun/ice-tea/internal/parser"
)

// TextParser returns minimal ParseResult (Source only, Root nil) for languages
// without AST support. Enables regex-only pattern matching.
type TextParser struct {
	languages []p.Language
}

// New creates a new Text parser for regex-only languages
func New() *TextParser {
	return &TextParser{
		languages: []p.Language{p.LangDart, p.LangZig, p.LangPerl},
	}
}

// SupportedLanguages returns the languages this parser handles
func (tp *TextParser) SupportedLanguages() []p.Language {
	return tp.languages
}

// SupportsLanguage checks if this parser handles the given language
func (tp *TextParser) SupportsLanguage(lang p.Language) bool {
	for _, l := range tp.languages {
		if l == lang {
			return true
		}
	}
	return false
}

// Parse returns a minimal ParseResult with Source only (Root nil).
// Regex patterns will match against Source; AST patterns will not match.
func (tp *TextParser) Parse(filename string, src []byte) (*p.ParseResult, error) {
	lang := detectLanguageFromFile(filename)
	if !tp.SupportsLanguage(lang) {
		return nil, &p.UnsupportedLanguageError{Language: lang}
	}

	return &p.ParseResult{
		Language: lang,
		Source:   src,
		FilePath: filename,
		Root:     nil,
	}, nil
}

func detectLanguageFromFile(filename string) p.Language {
	lower := strings.ToLower(filename)
	if strings.HasSuffix(lower, ".dart") {
		return p.LangDart
	}
	if strings.HasSuffix(lower, ".zig") {
		return p.LangZig
	}
	if strings.HasSuffix(lower, ".pl") || strings.HasSuffix(lower, ".pm") {
		return p.LangPerl
	}
	return ""
}
