package goImportVarShadowLint

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "goImportVarShadowLint",
	Doc:      "Check for variable names that conflict with package names",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	// 1. Build a set of imported-package names.
	imps := map[string]struct{}{}
	for _, imp := range pass.Pkg.Imports() {
		imps[imp.Name()] = struct{}{}
	}

	ins := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.GenDecl)(nil),
		(*ast.AssignStmt)(nil),
		(*ast.RangeStmt)(nil),
	}

	ins.Preorder(nodeFilter, func(n ast.Node) {
		var declIdents []*ast.Ident
		switch v := n.(type) {
		case *ast.AssignStmt:
			for _, lhs := range v.Lhs {
				if id, ok := lhs.(*ast.Ident); ok && id.Obj != nil {
					declIdents = append(declIdents, id)
				}
			}
		case *ast.ValueSpec:
			declIdents = v.Names
		case *ast.RangeStmt:
			if id, ok := v.Key.(*ast.Ident); ok && id.Obj != nil {
				declIdents = append(declIdents, id)
			}
			if id, ok := v.Value.(*ast.Ident); ok && id.Obj != nil {
				declIdents = append(declIdents, id)
			}
		case *ast.Field:
			declIdents = append(declIdents, v.Names...)
		}

		for _, decl := range declIdents {
			if decl == nil || decl.Name == "_" {
				continue
			}
			if _, clash := imps[decl.Name]; !clash {
				continue
			}

			obj := pass.TypesInfo.ObjectOf(decl)
			if obj == nil {
				continue
			}

			newName := decl.Name + "Var"
			var edits []analysis.TextEdit

			// ▲ 3. Find *all* idents that resolve to the same object.
			for _, f := range pass.Files {
				ast.Inspect(f, func(n ast.Node) bool {
					if id, ok := n.(*ast.Ident); ok {
						if pass.TypesInfo.ObjectOf(id) == obj {
							edits = append(edits, analysis.TextEdit{
								Pos:     id.Pos(),
								End:     id.End(),
								NewText: []byte(newName),
							})
						}
					}
					return true
				})
			}

			pass.Report(analysis.Diagnostic{
				Pos:     decl.Pos(),
				Message: fmt.Sprintf(`variable %q shadows imported package`, decl.Name),
				SuggestedFixes: []analysis.SuggestedFix{{
					Message:   fmt.Sprintf("rename %q → %q", decl.Name, newName),
					TextEdits: edits,
				}},
			})
		}
	})
	return nil, nil
}
