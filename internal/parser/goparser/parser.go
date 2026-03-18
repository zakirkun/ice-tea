package goparser

import (
	"fmt"
	"go/ast"
	goparser "go/parser"
	"go/token"

	p "github.com/zakirkun/ice-tea/internal/parser"
)

// GoParser parses Go source code using the standard library
type GoParser struct{}

// New creates a new Go parser
func New() *GoParser {
	return &GoParser{}
}

// SupportedLanguages returns the languages this parser handles
func (gp *GoParser) SupportedLanguages() []p.Language {
	return []p.Language{p.LangGo}
}

// SupportsLanguage checks if this parser handles the given language
func (gp *GoParser) SupportsLanguage(lang p.Language) bool {
	return lang == p.LangGo
}

// Parse parses Go source code into a common AST
func (gp *GoParser) Parse(filename string, src []byte) (*p.ParseResult, error) {
	fset := token.NewFileSet()
	file, err := goparser.ParseFile(fset, filename, src, goparser.AllErrors|goparser.ParseComments)

	result := &p.ParseResult{
		Language: p.LangGo,
		Source:   src,
		FilePath: filename,
	}

	if err != nil {
		// Collect parse errors but continue — partial AST may still be useful
		result.Errors = append(result.Errors, err.Error())
	}

	if file == nil {
		return result, fmt.Errorf("failed to parse Go file: %s", filename)
	}

	// Convert go/ast to our common AST format
	root := convertFile(fset, file, src)
	result.Root = root

	return result, nil
}

// convertFile converts a go/ast.File into our common Node format
func convertFile(fset *token.FileSet, file *ast.File, src []byte) *p.Node {
	root := &p.Node{
		Type:   "file",
		Fields: make(map[string]string),
	}

	if file.Name != nil {
		root.Fields["package"] = file.Name.Name
	}

	// Convert import declarations
	for _, imp := range file.Imports {
		importNode := convertImport(fset, imp)
		importNode.Parent = root
		root.Children = append(root.Children, importNode)
	}

	// Convert top-level declarations
	for _, decl := range file.Decls {
		nodes := convertDecl(fset, decl, src)
		for _, node := range nodes {
			node.Parent = root
			root.Children = append(root.Children, node)
		}
	}

	return root
}

// convertImport converts an import spec
func convertImport(fset *token.FileSet, imp *ast.ImportSpec) *p.Node {
	node := &p.Node{
		Type:   "import",
		Fields: make(map[string]string),
	}

	if imp.Path != nil {
		node.Fields["path"] = imp.Path.Value
		node.Text = imp.Path.Value
	}
	if imp.Name != nil {
		node.Fields["alias"] = imp.Name.Name
	}

	setPosition(fset, node, imp.Pos(), imp.End())
	return node
}

// convertDecl converts a top-level declaration
func convertDecl(fset *token.FileSet, decl ast.Decl, src []byte) []*p.Node {
	switch d := decl.(type) {
	case *ast.FuncDecl:
		return []*p.Node{convertFuncDecl(fset, d, src)}
	case *ast.GenDecl:
		return convertGenDecl(fset, d, src)
	default:
		return nil
	}
}

// convertFuncDecl converts a function declaration
func convertFuncDecl(fset *token.FileSet, fn *ast.FuncDecl, src []byte) *p.Node {
	node := &p.Node{
		Type:   "function_declaration",
		Fields: make(map[string]string),
	}

	if fn.Name != nil {
		node.Fields["name"] = fn.Name.Name
	}

	// Check if it's a method (has receiver)
	if fn.Recv != nil && fn.Recv.NumFields() > 0 {
		node.Type = "method_declaration"
		for _, field := range fn.Recv.List {
			node.Fields["receiver"] = typeString(field.Type)
		}
	}

	setPosition(fset, node, fn.Pos(), fn.End())
	node.Text = extractText(src, fset, fn.Pos(), fn.End())

	// Convert function body
	if fn.Body != nil {
		bodyNodes := convertBlockStmt(fset, fn.Body, src)
		for _, child := range bodyNodes {
			child.Parent = node
			node.Children = append(node.Children, child)
		}
	}

	return node
}

// convertGenDecl converts general declarations (var, const, type, import)
func convertGenDecl(fset *token.FileSet, gen *ast.GenDecl, src []byte) []*p.Node {
	var nodes []*p.Node

	for _, spec := range gen.Specs {
		switch s := spec.(type) {
		case *ast.ValueSpec:
			for _, name := range s.Names {
				node := &p.Node{
					Type:   "variable_declaration",
					Fields: map[string]string{"name": name.Name},
				}
				if gen.Tok == token.CONST {
					node.Type = "constant_declaration"
				}
				setPosition(fset, node, s.Pos(), s.End())
				node.Text = extractText(src, fset, s.Pos(), s.End())

				// Check for hardcoded string values
				for _, val := range s.Values {
					valNode := convertExpr(fset, val, src)
					if valNode != nil {
						valNode.Parent = node
						node.Children = append(node.Children, valNode)
					}
				}

				nodes = append(nodes, node)
			}
		case *ast.TypeSpec:
			node := &p.Node{
				Type:   "type_declaration",
				Fields: map[string]string{"name": s.Name.Name},
			}
			setPosition(fset, node, s.Pos(), s.End())
			node.Text = extractText(src, fset, s.Pos(), s.End())
			nodes = append(nodes, node)
		}
	}

	return nodes
}

// convertBlockStmt converts a block (list of statements)
func convertBlockStmt(fset *token.FileSet, block *ast.BlockStmt, src []byte) []*p.Node {
	if block == nil {
		return nil
	}

	var nodes []*p.Node
	for _, stmt := range block.List {
		stmtNodes := convertStmt(fset, stmt, src)
		nodes = append(nodes, stmtNodes...)
	}
	return nodes
}

// convertStmt converts a statement
func convertStmt(fset *token.FileSet, stmt ast.Stmt, src []byte) []*p.Node {
	switch s := stmt.(type) {
	case *ast.ExprStmt:
		if node := convertExpr(fset, s.X, src); node != nil {
			return []*p.Node{node}
		}
	case *ast.AssignStmt:
		return convertAssignStmt(fset, s, src)
	case *ast.ReturnStmt:
		node := &p.Node{
			Type:   "return_statement",
			Fields: make(map[string]string),
		}
		setPosition(fset, node, s.Pos(), s.End())
		node.Text = extractText(src, fset, s.Pos(), s.End())
		for _, expr := range s.Results {
			if child := convertExpr(fset, expr, src); child != nil {
				child.Parent = node
				node.Children = append(node.Children, child)
			}
		}
		return []*p.Node{node}
	case *ast.IfStmt:
		return []*p.Node{convertIfStmt(fset, s, src)}
	case *ast.ForStmt:
		return convertForStmt(fset, s, src)
	case *ast.RangeStmt:
		return convertRangeStmt(fset, s, src)
	case *ast.DeclStmt:
		if gen, ok := s.Decl.(*ast.GenDecl); ok {
			return convertGenDecl(fset, gen, src)
		}
	case *ast.DeferStmt:
		node := &p.Node{
			Type:   "defer_statement",
			Fields: make(map[string]string),
		}
		setPosition(fset, node, s.Pos(), s.End())
		if callNode := convertExpr(fset, s.Call, src); callNode != nil {
			callNode.Parent = node
			node.Children = append(node.Children, callNode)
		}
		return []*p.Node{node}
	case *ast.GoStmt:
		node := &p.Node{
			Type:   "go_statement",
			Fields: make(map[string]string),
		}
		setPosition(fset, node, s.Pos(), s.End())
		if callNode := convertExpr(fset, s.Call, src); callNode != nil {
			callNode.Parent = node
			node.Children = append(node.Children, callNode)
		}
		return []*p.Node{node}
	case *ast.BlockStmt:
		return convertBlockStmt(fset, s, src)
	}
	return nil
}

// convertAssignStmt converts an assignment statement
func convertAssignStmt(fset *token.FileSet, assign *ast.AssignStmt, src []byte) []*p.Node {
	node := &p.Node{
		Type:   "assignment",
		Fields: map[string]string{"operator": assign.Tok.String()},
	}
	setPosition(fset, node, assign.Pos(), assign.End())
	node.Text = extractText(src, fset, assign.Pos(), assign.End())

	for _, rhs := range assign.Rhs {
		if child := convertExpr(fset, rhs, src); child != nil {
			child.Parent = node
			node.Children = append(node.Children, child)
		}
	}
	return []*p.Node{node}
}

// convertExpr converts an expression
func convertExpr(fset *token.FileSet, expr ast.Expr, src []byte) *p.Node {
	switch e := expr.(type) {
	case *ast.CallExpr:
		return convertCallExpr(fset, e, src)
	case *ast.BasicLit:
		node := &p.Node{
			Type:   "literal",
			Fields: map[string]string{"kind": e.Kind.String()},
			Text:   e.Value,
		}
		setPosition(fset, node, e.Pos(), e.End())
		return node
	case *ast.BinaryExpr:
		node := &p.Node{
			Type:   "binary_expression",
			Fields: map[string]string{"operator": e.Op.String()},
		}
		setPosition(fset, node, e.Pos(), e.End())
		node.Text = extractText(src, fset, e.Pos(), e.End())
		if left := convertExpr(fset, e.X, src); left != nil {
			left.Parent = node
			node.Children = append(node.Children, left)
		}
		if right := convertExpr(fset, e.Y, src); right != nil {
			right.Parent = node
			node.Children = append(node.Children, right)
		}
		return node
	case *ast.Ident:
		return &p.Node{
			Type:   "identifier",
			Text:   e.Name,
			Fields: map[string]string{"name": e.Name},
			Start:  posToPosition(fset, e.Pos()),
			End:    posToPosition(fset, e.End()),
		}
	case *ast.SelectorExpr:
		node := &p.Node{
			Type:   "selector_expression",
			Fields: make(map[string]string),
		}
		setPosition(fset, node, e.Pos(), e.End())
		node.Text = extractText(src, fset, e.Pos(), e.End())
		if e.Sel != nil {
			node.Fields["selector"] = e.Sel.Name
		}
		if ident, ok := e.X.(*ast.Ident); ok {
			node.Fields["object"] = ident.Name
		}
		return node
	case *ast.CompositeLit:
		node := &p.Node{
			Type:   "composite_literal",
			Fields: make(map[string]string),
		}
		setPosition(fset, node, e.Pos(), e.End())
		node.Text = extractText(src, fset, e.Pos(), e.End())
		return node
	case *ast.FuncLit:
		node := &p.Node{
			Type:   "function_literal",
			Fields: make(map[string]string),
		}
		setPosition(fset, node, e.Pos(), e.End())
		if e.Body != nil {
			bodyNodes := convertBlockStmt(fset, e.Body, src)
			for _, child := range bodyNodes {
				child.Parent = node
				node.Children = append(node.Children, child)
			}
		}
		return node
	}
	return nil
}

// convertCallExpr converts a call expression
func convertCallExpr(fset *token.FileSet, call *ast.CallExpr, src []byte) *p.Node {
	node := &p.Node{
		Type:   "call_expression",
		Fields: make(map[string]string),
	}
	setPosition(fset, node, call.Pos(), call.End())
	node.Text = extractText(src, fset, call.Pos(), call.End())

	// Extract function name
	switch fn := call.Fun.(type) {
	case *ast.Ident:
		node.Fields["function"] = fn.Name
	case *ast.SelectorExpr:
		if ident, ok := fn.X.(*ast.Ident); ok {
			node.Fields["function"] = ident.Name + "." + fn.Sel.Name
			node.Fields["object"] = ident.Name
			node.Fields["method"] = fn.Sel.Name
		} else {
			node.Fields["function"] = fn.Sel.Name
			node.Fields["method"] = fn.Sel.Name
			if childNode := convertExpr(fset, fn.X, src); childNode != nil {
				childNode.Parent = node
				node.Children = append(node.Children, childNode)
			}
		}
	}

	// Convert arguments
	for _, arg := range call.Args {
		if argNode := convertExpr(fset, arg, src); argNode != nil {
			argNode.Parent = node
			node.Children = append(node.Children, argNode)
		}
	}

	return node
}

// convertIfStmt converts an if statement
func convertIfStmt(fset *token.FileSet, ifStmt *ast.IfStmt, src []byte) *p.Node {
	node := &p.Node{
		Type:   "if_statement",
		Fields: make(map[string]string),
	}
	setPosition(fset, node, ifStmt.Pos(), ifStmt.End())

	if ifStmt.Body != nil {
		bodyNodes := convertBlockStmt(fset, ifStmt.Body, src)
		for _, child := range bodyNodes {
			child.Parent = node
			node.Children = append(node.Children, child)
		}
	}
	return node
}

// convertForStmt converts a for statement
func convertForStmt(fset *token.FileSet, forStmt *ast.ForStmt, src []byte) []*p.Node {
	node := &p.Node{
		Type:   "for_statement",
		Fields: make(map[string]string),
	}
	setPosition(fset, node, forStmt.Pos(), forStmt.End())

	if forStmt.Body != nil {
		bodyNodes := convertBlockStmt(fset, forStmt.Body, src)
		for _, child := range bodyNodes {
			child.Parent = node
			node.Children = append(node.Children, child)
		}
	}
	return []*p.Node{node}
}

// convertRangeStmt converts a range statement
func convertRangeStmt(fset *token.FileSet, rangeStmt *ast.RangeStmt, src []byte) []*p.Node {
	node := &p.Node{
		Type:   "range_statement",
		Fields: make(map[string]string),
	}
	setPosition(fset, node, rangeStmt.Pos(), rangeStmt.End())

	if rangeStmt.Body != nil {
		bodyNodes := convertBlockStmt(fset, rangeStmt.Body, src)
		for _, child := range bodyNodes {
			child.Parent = node
			node.Children = append(node.Children, child)
		}
	}
	return []*p.Node{node}
}

// Helper functions

func setPosition(fset *token.FileSet, node *p.Node, start, end token.Pos) {
	node.Start = posToPosition(fset, start)
	node.End = posToPosition(fset, end)
}

func posToPosition(fset *token.FileSet, pos token.Pos) p.Position {
	if !pos.IsValid() {
		return p.Position{}
	}
	position := fset.Position(pos)
	return p.Position{
		Line:   position.Line,
		Column: position.Column,
		Offset: position.Offset,
	}
}

func extractText(src []byte, fset *token.FileSet, start, end token.Pos) string {
	if !start.IsValid() || !end.IsValid() {
		return ""
	}
	startPos := fset.Position(start)
	endPos := fset.Position(end)
	if startPos.Offset >= 0 && endPos.Offset <= len(src) && startPos.Offset < endPos.Offset {
		return string(src[startPos.Offset:endPos.Offset])
	}
	return ""
}

func typeString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return "*" + typeString(t.X)
	case *ast.SelectorExpr:
		return typeString(t.X) + "." + t.Sel.Name
	default:
		return ""
	}
}
