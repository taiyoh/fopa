package fopa

import (
	"fmt"
	"go/ast"
	"sort"
	"strings"
)

type syntaxTree struct {
	name    string
	imports map[string]importPath
	fields  []field
}

func findAst(base string, astf *ast.File) (*syntaxTree, error) {
	imports := findImports(astf)
	baseTyp, exists := astf.Scope.Objects[base]
	if !exists {
		return nil, fmt.Errorf("type:%s not found", base)
	}

	typeSpec := baseTyp.Decl.(*ast.TypeSpec).Type.(*ast.StructType)
	fields := []field{}
	for _, f := range typeSpec.Fields.List {
		fields = append(fields, newField(f))
	}

	return &syntaxTree{baseTyp.Name, imports, fields}, nil
}

func (s *syntaxTree) Name() string {
	return s.name
}

func (s *syntaxTree) ImportPaths() []importPath {
	imports := []importPath{}
	uniq := map[string]struct{}{}
	for _, f := range s.fields {
		if f.importPkg == "" {
			continue
		}
		if _, ok := uniq[f.importPkg]; ok {
			continue
		}
		imports = append(imports, s.imports[f.importPkg])
		uniq[f.importPkg] = struct{}{}
	}
	sort.Slice(imports, func(i, j int) bool {
		return imports[i].order < imports[j].order
	})
	return imports
}

func findImports(astf *ast.File) map[string]importPath {
	imports := map[string]importPath{}
	for i, spec := range astf.Imports {
		path := trimPath(spec.Path.Value)
		sig, exported := findSig(path, spec)
		imports[sig] = importPath{
			order:    i,
			sig:      sig,
			path:     path,
			exported: exported,
		}
	}
	return imports
}

func trimPath(p string) string {
	val := []rune(p)
	if val[0] == rune(34) {
		val = val[1:]
	}
	if val[len(val)-1] == rune(34) {
		val = val[:len(val)-1]
	}
	return string(val)
}

func findSig(path string, spec *ast.ImportSpec) (string, bool) {
	if spec.Name != nil {
		return spec.Name.Name, true
	}
	parts := strings.Split(path, "/")
	sig := parts[len(parts)-1]
	return sig, false
}
