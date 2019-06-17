package fopa

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"unicode"
)

func findFile(baseDir, filename string) (*File, error) {
	found, dir := findpath(baseDir, filename)
	if found == "" {
		return nil, nil
	}
	data, err := ioutil.ReadFile(found)
	if err != nil {
		return nil, err
	}
	return &File{dir, filename, data}, nil
}

func findpath(dir, target string) (string, string) {
	infolist, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", ""
	}
	subdirs := []string{}
	for _, fileinfo := range infolist {
		name := fileinfo.Name()
		p := filepath.Join(dir, name)
		if fileinfo.IsDir() {
			subdirs = append(subdirs, p)
			continue
		}
		if name == target {
			return p, dir
		}
	}
	for _, d := range subdirs {
		if found, dir := findpath(d, target); found != "" {
			return found, dir
		}
	}
	return "", ""
}

func buildFile(ast *syntaxTree, baseDir, pkgname, factory, builder string) *File {
	b := bytes.NewBuffer([]byte{})
	fmt.Fprintf(b, "// Code generated by fopa. DO NOT EDIT.\n\n")
	fmt.Fprintf(b, "package %s\n\n", pkgname)

	if imports := ast.ImportPaths(); len(imports) > 0 {
		fmt.Fprint(b, "import (\n")
		for _, i := range imports {
			fmt.Fprintf(b, "\t%s\n", i.String())
		}
		fmt.Fprint(b, ")\n\n")
	}

	fmt.Fprintf(b, "type %s func(*%s)\n\n", builder, ast.Name())

	objName := ast.Name()
	fmt.Fprintf(b, "func (f *%s) Setup(fns ...%s) *%s {\n", factory, builder, objName)
	fmt.Fprintf(b, "\to := &%s{}\n", objName)
	fmt.Fprint(b, "\tfor _, fn := range fns {\n")
	fmt.Fprint(b, "\t\tfn(o)\n")
	fmt.Fprint(b, "\t}\n")
	fmt.Fprint(b, "\treturn o\n")
	fmt.Fprint(b, "}\n")

	for _, f := range ast.fields {
		if !f.tag.enabled {
			continue
		}
		fmt.Fprintf(b, "\nfunc (f *%s) Fill%s(%s) %s {\n", factory, f.titleName(), f.args(), builder)
		fmt.Fprintf(b, "\treturn func(p *%s) {\n", ast.Name())
		fmt.Fprintf(b, "\t\tp.%s = %s\n", f.name, f.expr())
		fmt.Fprintf(b, "\t}\n")
		fmt.Fprintf(b, "}\n")
	}

	return &File{
		baseDir:  baseDir,
		filename: fmt.Sprintf("%s_gen.go", toSnake(factory)),
		data:     b.Bytes(),
	}
}

func toSnake(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}
