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
	var typName string
	switch f.Type.(type) {
	case *ast.Ident:
		typ := f.Type.(*ast.Ident)
		typName = typ.Name
	case *ast.SelectorExpr:
		typ := f.Type.(*ast.SelectorExpr)
		typBase := typ.X.(*ast.Ident)
		typName = fmt.Sprintf("%s.%s", typBase.Name, typ.Sel.Name)
	}
	return field{
		name: f.Names[0].Name,
		typ:  typName,
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
	if v, exists := commonInitialisms[f.name]; exists {
		return v
	}
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

// this mapping from https://github.com/golang/lint/blob/master/lint.go
var commonInitialisms = map[string]string{
	"acl":   "ACL",
	"api":   "API",
	"ascii": "ASCII",
	"cpu":   "CPU",
	"css":   "CSS",
	"dns":   "DNS",
	"eof":   "EOF",
	"guid":  "GUID",
	"html":  "HTML",
	"http":  "HTTP",
	"https": "HTTPS",
	"id":    "ID",
	"ip":    "IP",
	"json":  "JSON",
	"lhs":   "LHS",
	"qps":   "QPS",
	"ram":   "RAM",
	"rhs":   "RHS",
	"rpc":   "RPC",
	"sla":   "SLA",
	"smtp":  "SMTP",
	"sql":   "SQL",
	"ssh":   "SSH",
	"tcp":   "TCP",
	"tls":   "TLS",
	"ttl":   "TTL",
	"udp":   "UDP",
	"ui":    "UI",
	"uid":   "UID",
	"uuid":  "UUID",
	"uri":   "URI",
	"url":   "URL",
	"utf8":  "UTF8",
	"vm":    "VM",
	"xml":   "XML",
	"xmpp":  "XMPP",
	"xsrf":  "XSRF",
	"xss":   "XSS",
}
