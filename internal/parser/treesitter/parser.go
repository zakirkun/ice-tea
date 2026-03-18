package treesitter

import (
	"context"
	"fmt"
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/c"
	"github.com/smacker/go-tree-sitter/cpp"
	"github.com/smacker/go-tree-sitter/java"
	"github.com/smacker/go-tree-sitter/javascript"
	"github.com/smacker/go-tree-sitter/php"
	"github.com/smacker/go-tree-sitter/python"
	"github.com/smacker/go-tree-sitter/ruby"
	"github.com/smacker/go-tree-sitter/rust"
	"github.com/smacker/go-tree-sitter/typescript/typescript"

	p "github.com/zakirkun/ice-tea/internal/parser"
)

// TreeSitterParser parses source code using Tree-Sitter grammars.
type TreeSitterParser struct {
	languages map[p.Language]*sitter.Language
}

// New creates a new Tree-Sitter parser
func New() *TreeSitterParser {
	return &TreeSitterParser{
		languages: map[p.Language]*sitter.Language{
			p.LangJavaScript: javascript.GetLanguage(),
			p.LangTypeScript: typescript.GetLanguage(),
			p.LangPython:     python.GetLanguage(),
			p.LangJava:       java.GetLanguage(),
			p.LangPHP:        php.GetLanguage(),
			p.LangRuby:       ruby.GetLanguage(),
			p.LangRust:       rust.GetLanguage(),
			p.LangC:          c.GetLanguage(),
			p.LangCPP:        cpp.GetLanguage(),
		},
	}
}

// SupportedLanguages returns the languages this parser handles
func (ts *TreeSitterParser) SupportedLanguages() []p.Language {
	langs := make([]p.Language, 0, len(ts.languages))
	for lang := range ts.languages {
		langs = append(langs, lang)
	}
	return langs
}

// SupportsLanguage checks if this parser handles the given language
func (ts *TreeSitterParser) SupportsLanguage(lang p.Language) bool {
	_, ok := ts.languages[lang]
	return ok
}

// Parse parses source code into a common AST using Tree-Sitter.
func (ts *TreeSitterParser) Parse(filename string, src []byte) (*p.ParseResult, error) {
	lang := detectLanguageFromFile(filename)
	tsLang, ok := ts.languages[lang]
	if !ok {
		return nil, fmt.Errorf("unsupported language for tree-sitter: %s", lang)
	}

	parser := sitter.NewParser()
	parser.SetLanguage(tsLang)

	tree, err := parser.ParseCtx(context.Background(), nil, src)
	if err != nil {
		return nil, fmt.Errorf("tree-sitter parsing failed: %w", err)
	}

	result := &p.ParseResult{
		Language: lang,
		Source:   src,
		FilePath: filename,
	}

	rootTS := tree.RootNode()
	result.Root = convertSitterNode(rootTS, src)

	return result, nil
}

func convertSitterNode(tsNode *sitter.Node, src []byte) *p.Node {
	if tsNode == nil {
		return nil
	}

	start := tsNode.StartPoint()
	end := tsNode.EndPoint()

	node := &p.Node{
		Type: tsNode.Type(),
		Text: tsNode.Content(src),
		Start: p.Position{
			Line:   int(start.Row + 1),
			Column: int(start.Column),
			Offset: int(tsNode.StartByte()),
		},
		End: p.Position{
			Line:   int(end.Row + 1),
			Column: int(end.Column),
			Offset: int(tsNode.EndByte()),
		},
		Fields: make(map[string]string),
	}

	count := int(tsNode.ChildCount())
	for i := 0; i < count; i++ {
		childTS := tsNode.Child(i)

		// Map field names
		fieldName := tsNode.FieldNameForChild(i)
		if fieldName != "" {
			node.Fields[fieldName] = childTS.Content(src)
		}

		childNode := convertSitterNode(childTS, src)
		if childNode != nil {
			childNode.Parent = node
			node.Children = append(node.Children, childNode)
		}
	}

	return node
}

// detectLanguageFromFile determines language from file extension
func detectLanguageFromFile(filename string) p.Language {
	if strings.HasSuffix(filename, ".js") || strings.HasSuffix(filename, ".jsx") {
		return p.LangJavaScript
	}
	if strings.HasSuffix(filename, ".ts") || strings.HasSuffix(filename, ".tsx") {
		return p.LangTypeScript
	}
	if strings.HasSuffix(filename, ".py") {
		return p.LangPython
	}
	if strings.HasSuffix(filename, ".java") {
		return p.LangJava
	}
	if strings.HasSuffix(filename, ".php") {
		return p.LangPHP
	}
	if strings.HasSuffix(filename, ".rb") {
		return p.LangRuby
	}
	if strings.HasSuffix(filename, ".rs") {
		return p.LangRust
	}
	if strings.HasSuffix(filename, ".c") || strings.HasSuffix(filename, ".h") {
		return p.LangC
	}
	if strings.HasSuffix(filename, ".cpp") || strings.HasSuffix(filename, ".hpp") || strings.HasSuffix(filename, ".cc") {
		return p.LangCPP
	}
	return "unknown"
}
