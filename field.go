package fopa

import (
	"go/ast"
	"reflect"
)

type field struct {
	name string
	typ  string
	tag  tag
}

func newField(f *ast.Field) field {
	rawtag := trimTag(f.Tag.Value)
	tagVal := reflect.StructTag(rawtag).Get("fopa")
	typ := f.Type.(*ast.Ident)
	return field{
		name: f.Names[0].Name,
		typ:  typ.Name,
		tag:  tag{tagVal},
	}
}

func trimTag(tag string) string {
	runes := []rune(tag)
	if runes[0] == rune(96) {
		runes = runes[1:]
	}
	last := len(runes) - 1
	if runes[last] == rune(96) {
		runes = runes[:len(runes)-1]
	}
	return string(runes)
}
