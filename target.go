package fopa

import (
	"fmt"
	"go/ast"
	"go/parser"
	"io/ioutil"

	"golang.org/x/tools/go/loader"
)

type Target struct {
	name   string
	fields []field
}

func FindTarget(base, basedir, filename string) (*Target, error) {
	found := findpath(basedir, filename)
	if found == "" {
		return nil, nil
	}
	data, err := ioutil.ReadFile(found)
	if err != nil {
		return nil, err
	}
	loader := loader.Config{ParserMode: parser.ParseComments}
	astf, err := loader.ParseFile(filename, string(data))
	if err != nil {
		return nil, err
	}
	baseTyp, exists := astf.Scope.Objects[base]
	if !exists {
		return nil, fmt.Errorf("type:%s not found", base)
	}

	typeSpec := baseTyp.Decl.(*ast.TypeSpec).Type.(*ast.StructType)
	fields := []field{}
	for _, f := range typeSpec.Fields.List {
		fields = append(fields, newField(f))
	}

	return &Target{
		name:   baseTyp.Name,
		fields: fields,
	}, nil
}
