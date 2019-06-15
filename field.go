package fopa

import (
	"fmt"
	"go/ast"
	"reflect"
	"strings"
)

type field struct {
	name string
	typ  string
	tag  tag
}

func newField(f *ast.Field) field {
	var tagVal string
	if t := f.Tag; t != nil {
		rawtag := trimTag(f.Tag.Value)
		tagVal = reflect.StructTag(rawtag).Get("fopa")
	}
	typ := f.Type.(*ast.Ident)
	return field{
		name: f.Names[0].Name,
		typ:  typ.Name,
		tag:  newTag(tagVal),
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

func (f field) titleName() string {
	return strings.Title(f.name)
}

func (f field) args() string {
	t := f.tag
	if t.acceptType == "" {
		return fmt.Sprintf("v %s", f.typ)
	}
	if t.acceptArgCount() == 1 {
		return fmt.Sprintf("v %s", t.acceptType)
	}
	return ""
}

func (f field) expr() string {
	t := f.tag
	if t.fillExpr == "" {
		if t.acceptType == "" || t.acceptType == f.typ {
			return "v"
		}
		return fmt.Sprintf("%s(v)", f.typ)
	}
	if t.exprStepsCount() == 1 {
		return strings.Replace(t.fillExpr, "{}", "v", 1)
	}
	return ""
}
