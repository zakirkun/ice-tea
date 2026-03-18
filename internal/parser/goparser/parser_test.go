package goparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	p "github.com/zakirkun/ice-tea/internal/parser"
)

func TestGoParserSupportsLanguage(t *testing.T) {
	gp := New()
	assert.True(t, gp.SupportsLanguage(p.LangGo))
	assert.False(t, gp.SupportsLanguage(p.LangJavaScript))
	assert.False(t, gp.SupportsLanguage(p.LangPython))
}

func TestGoParserSupportedLanguages(t *testing.T) {
	gp := New()
	langs := gp.SupportedLanguages()
	assert.Equal(t, []p.Language{p.LangGo}, langs)
}

func TestParseSimpleFunction(t *testing.T) {
	src := []byte(`package main

import "fmt"

func main() {
	fmt.Println("hello")
}
`)

	gp := New()
	result, err := gp.Parse("main.go", src)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, p.LangGo, result.Language)
	assert.Equal(t, "main.go", result.FilePath)
	assert.NotNil(t, result.Root)
	assert.Equal(t, "file", result.Root.Type)
	assert.Equal(t, "main", result.Root.Fields["package"])

	// Should have import and function
	assert.True(t, len(result.Root.Children) >= 2)

	// Find the function declaration
	funcs := p.FindAll(result.Root, "function_declaration")
	assert.Len(t, funcs, 1)
	assert.Equal(t, "main", funcs[0].Fields["name"])

	// Find the call expression
	calls := p.FindAll(result.Root, "call_expression")
	assert.True(t, len(calls) >= 1)
	assert.Equal(t, "fmt.Println", calls[0].Fields["function"])
}

func TestParseCallExpressions(t *testing.T) {
	src := []byte(`package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	query := "SELECT * FROM users WHERE id = " + id
	db.Query(query)
	fmt.Fprintf(w, "Result")
}
`)

	gp := New()
	result, err := gp.Parse("handler.go", src)
	require.NoError(t, err)
	require.NotNil(t, result)

	// Find call expressions
	calls := p.FindAll(result.Root, "call_expression")
	assert.True(t, len(calls) >= 2, "should find multiple call expressions")

	// Find binary expressions (string concatenation)
	binaries := p.FindAll(result.Root, "binary_expression")
	found := false
	for _, b := range binaries {
		if b.Fields["operator"] == "+" {
			found = true
			break
		}
	}
	assert.True(t, found, "should find string concatenation")
}

func TestParseMethodDeclaration(t *testing.T) {
	src := []byte(`package main

type Server struct {
	port int
}

func (s *Server) Start() error {
	return nil
}
`)

	gp := New()
	result, err := gp.Parse("server.go", src)
	require.NoError(t, err)

	methods := p.FindAll(result.Root, "method_declaration")
	assert.Len(t, methods, 1)
	assert.Equal(t, "Start", methods[0].Fields["name"])
	assert.Equal(t, "*Server", methods[0].Fields["receiver"])
}

func TestParseImports(t *testing.T) {
	src := []byte(`package main

import (
	"crypto/md5"
	"fmt"
	"os/exec"
)

func main() {}
`)

	gp := New()
	result, err := gp.Parse("main.go", src)
	require.NoError(t, err)

	imports := p.FindAll(result.Root, "import")
	assert.Len(t, imports, 3)

	importPaths := make([]string, 0)
	for _, imp := range imports {
		importPaths = append(importPaths, imp.Fields["path"])
	}
	assert.Contains(t, importPaths, "\"crypto/md5\"")
	assert.Contains(t, importPaths, "\"os/exec\"")
}

func TestParseInvalidCode(t *testing.T) {
	src := []byte(`package main

func broken( {
	// invalid syntax
`)

	gp := New()
	result, err := gp.Parse("broken.go", src)
	// Should still return a result even with errors
	require.NoError(t, err) // err is only for total failure
	require.NotNil(t, result)
	assert.True(t, len(result.Errors) > 0, "should have parse errors")
}

func TestFindAll(t *testing.T) {
	root := &p.Node{
		Type: "file",
		Children: []*p.Node{
			{
				Type: "function_declaration",
				Fields: map[string]string{"name": "foo"},
				Children: []*p.Node{
					{Type: "call_expression", Fields: map[string]string{"function": "bar"}},
					{Type: "call_expression", Fields: map[string]string{"function": "baz"}},
				},
			},
			{
				Type: "function_declaration",
				Fields: map[string]string{"name": "qux"},
			},
		},
	}

	funcs := p.FindAll(root, "function_declaration")
	assert.Len(t, funcs, 2)

	calls := p.FindAll(root, "call_expression")
	assert.Len(t, calls, 2)
}
