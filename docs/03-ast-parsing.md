# AST Parsing & Code Analysis

## Overview

Ice Tea analyzes source code by parsing it into **Abstract Syntax Trees (ASTs)** and then traversing those trees to find vulnerability patterns. Two parsing backends are used:

1. **Go native parser** (`go/parser` + `go/ast`) — for Go source code
2. **Tree-Sitter** — for all other supported languages

## Go Native AST Parsing

Go's standard library includes robust packages for parsing Go source code:

### Key Packages

| Package | Purpose |
|---------|---------|
| `go/token` | Defines token types and file position tracking |
| `go/scanner` | Lexical scanning (tokenization) |
| `go/parser` | Parses Go source into AST nodes |
| `go/ast` | Defines AST node types and traversal utilities |

### Usage Pattern

```go
import (
    "go/ast"
    "go/parser"
    "go/token"
)

func parseGoFile(filename string, src []byte) (*ast.File, error) {
    fset := token.NewFileSet()
    file, err := parser.ParseFile(fset, filename, src, parser.AllErrors)
    if err != nil {
        return nil, err
    }
    return file, nil
}

// Traverse the AST to find dangerous patterns
func inspectForVulnerabilities(file *ast.File) []Finding {
    var findings []Finding

    ast.Inspect(file, func(n ast.Node) bool {
        // Example: detect calls to exec.Command with variables
        call, ok := n.(*ast.CallExpr)
        if !ok {
            return true
        }
        // Check if calling dangerous functions...
        // Add finding if pattern matches
        return true
    })

    return findings
}
```

### Advantages
- **Zero dependencies** — part of Go's standard library
- **Exact representation** — full fidelity of Go syntax
- **SSA support** — `golang.org/x/tools/go/ssa` enables data flow analysis
- **Type information** — `go/types` provides complete type checking

## Tree-Sitter Multi-Language Parsing

Tree-Sitter is an incremental parsing library designed for:
- **Speed**: Parses in real-time, suitable for large codebases
- **Robustness**: Produces useful ASTs even for syntactically invalid code
- **Multi-language**: Supports 100+ programming languages via grammar plugins
- **Dependency-free runtime**: No external dependencies at parse time

### Go Bindings

```go
import (
    ts "github.com/tree-sitter/go-tree-sitter"
    javascript "github.com/tree-sitter/tree-sitter-javascript/bindings/go"
    python "github.com/tree-sitter/tree-sitter-python/bindings/go"
)

func parseJavaScript(code []byte) *ts.Tree {
    parser := ts.NewParser()
    defer parser.Close()

    parser.SetLanguage(ts.NewLanguage(javascript.Language()))
    tree := parser.Parse(code, nil)
    return tree
}

func parsePython(code []byte) *ts.Tree {
    parser := ts.NewParser()
    defer parser.Close()

    parser.SetLanguage(ts.NewLanguage(python.Language()))
    tree := parser.Parse(code, nil)
    return tree
}
```

### Supported Languages (Initial Set)

| Language | Grammar Package | Priority |
|----------|----------------|----------|
| Go | `go/parser` (native) | P0 |
| JavaScript/TypeScript | `tree-sitter-javascript`, `tree-sitter-typescript` | P0 |
| Python | `tree-sitter-python` | P0 |
| Java | `tree-sitter-java` | P1 |
| PHP | `tree-sitter-php` | P1 |
| Ruby | `tree-sitter-ruby` | P2 |
| Rust | `tree-sitter-rust` | P2 |
| C/C++ | `tree-sitter-c`, `tree-sitter-cpp` | P2 |

### Tree-Sitter Query Language

Tree-Sitter includes a powerful **query language** for pattern matching on ASTs. This is ideal for defining vulnerability detection rules:

```scheme
; Find all function calls where the callee is "exec"
(call_expression
  function: (identifier) @func-name
  (#eq? @func-name "exec")
  arguments: (arguments) @args
) @dangerous-call

; Find string concatenation in SQL-like contexts
(binary_expression
  left: (string) @sql-string
  operator: "+"
  right: (_) @concatenated
  (#match? @sql-string "SELECT|INSERT|UPDATE|DELETE|WHERE")
) @sql-concat
```

## Common AST Interface

Ice Tea defines a common interface that abstracts away the parser backend:

```go
// Parser interface for language-agnostic AST handling
type Parser interface {
    // Parse source code into an AST
    Parse(filename string, src []byte) (AST, error)
    // SupportsLanguage checks if this parser handles the given language
    SupportsLanguage(lang Language) bool
}

// AST is a common interface over parsed code
type AST interface {
    // Walk traverses the AST, calling the visitor for each node
    Walk(visitor Visitor) error
    // Query runs a pattern query against the AST
    Query(pattern string) ([]Match, error)
    // Source returns the original source code
    Source() []byte
}

// Visitor is called for each node during AST traversal
type Visitor func(node Node) WalkAction

// Node represents a generic AST node
type Node interface {
    Type() string           // Node type name
    StartPosition() Position
    EndPosition() Position
    Children() []Node
    Text() string           // Source text of this node
    Parent() Node
}
```

## Taint Analysis

Beyond pattern matching, Ice Tea performs **taint analysis** to track data flow:

```
Source (untrusted input) ──► Propagation (through variables/functions) ──► Sink (dangerous operation)
```

### Taint Sources (Examples)
- HTTP request parameters, headers, body
- File reads
- Database query results
- Environment variables (in some contexts)
- Command-line arguments

### Taint Sinks (Examples)
- SQL query execution (`db.Query`, `db.Exec`)
- Command execution (`exec.Command`, `os.system`)
- File operations (`os.Create`, `os.Open` with user paths)
- HTML rendering without escaping
- Deserialization of untrusted data

### Analysis Flow
1. Identify taint sources in the AST
2. Track variable assignments and function calls
3. Check if tainted data reaches a sink without sanitization
4. Report findings with full data flow path
