package parser

// Language represents a programming language identifier
type Language string

const (
	LangGo         Language = "go"
	LangJavaScript Language = "javascript"
	LangTypeScript Language = "typescript"
	LangPython     Language = "python"
	LangJava       Language = "java"
	LangPHP        Language = "php"
	LangRuby       Language = "ruby"
	LangRust       Language = "rust"
	LangC          Language = "c"
	LangCPP        Language = "cpp"
	LangShell      Language = "shell"
	LangYAML       Language = "yaml"
	LangDart       Language = "dart"
	LangKotlin     Language = "kotlin"
	LangZig        Language = "zig"
	LangPerl       Language = "perl"
	LangElixir     Language = "elixir"
)

// Position represents a position in source code
type Position struct {
	Line   int // 1-based line number
	Column int // 0-based column offset
	Offset int // byte offset from start of file
}

// Node represents a generic AST node
type Node struct {
	Type     string   // node type name (e.g., "function_declaration", "call_expression")
	Text     string   // source text of this node
	Start    Position // start position
	End      Position // end position
	Children []*Node  // child nodes
	Parent   *Node    // parent node (nil for root)
	Fields   map[string]string // named fields (e.g., "name", "value")
}

// ParseResult holds the result of parsing a file
type ParseResult struct {
	Language Language  // detected language
	Root     *Node     // root AST node
	Source   []byte    // original source code
	FilePath string    // file path
	Errors   []string  // parse errors (non-fatal)
}

// Parser is the common interface for all AST parsers
type Parser interface {
	// Parse parses source code and returns an AST
	Parse(filename string, src []byte) (*ParseResult, error)

	// SupportedLanguages returns the languages this parser handles
	SupportedLanguages() []Language

	// SupportsLanguage checks if this parser handles the given language
	SupportsLanguage(lang Language) bool
}

// WalkFunc is called for each node during AST traversal
type WalkFunc func(node *Node) bool

// Walk traverses the AST in depth-first order.
// The walkFunc is called for each node. If it returns false, the
// children of that node are not visited.
func Walk(node *Node, fn WalkFunc) {
	if node == nil {
		return
	}
	if !fn(node) {
		return
	}
	for _, child := range node.Children {
		Walk(child, fn)
	}
}

// FindAll returns all nodes matching the given type
func FindAll(root *Node, nodeType string) []*Node {
	var matches []*Node
	Walk(root, func(n *Node) bool {
		if n.Type == nodeType {
			matches = append(matches, n)
		}
		return true
	})
	return matches
}

// FindByText returns all nodes whose text matches the given string
func FindByText(root *Node, text string) []*Node {
	var matches []*Node
	Walk(root, func(n *Node) bool {
		if n.Text == text {
			matches = append(matches, n)
		}
		return true
	})
	return matches
}

// Registry holds registered parsers
type Registry struct {
	parsers []Parser
}

// NewRegistry creates a new parser registry
func NewRegistry() *Registry {
	return &Registry{}
}

// Register adds a parser to the registry
func (r *Registry) Register(p Parser) {
	r.parsers = append(r.parsers, p)
}

// GetParser returns a parser that supports the given language
func (r *Registry) GetParser(lang Language) Parser {
	for _, p := range r.parsers {
		if p.SupportsLanguage(lang) {
			return p
		}
	}
	return nil
}

// Parse parses a file using the appropriate parser for its language
func (r *Registry) Parse(filename string, src []byte, lang Language) (*ParseResult, error) {
	p := r.GetParser(lang)
	if p == nil {
		return nil, &UnsupportedLanguageError{Language: lang}
	}
	return p.Parse(filename, src)
}

// UnsupportedLanguageError is returned when no parser supports a language
type UnsupportedLanguageError struct {
	Language Language
}

func (e *UnsupportedLanguageError) Error() string {
	return "no parser available for language: " + string(e.Language)
}
