package fopa

import (
	"go/ast"
	"go/parser"
	"io/ioutil"
	"path/filepath"

	"golang.org/x/tools/go/loader"
)

type file struct {
	baseDir  string
	filename string
	data     []byte
}

func findFile(baseDir, filename string) (*file, error) {
	found := findpath(baseDir, filename)
	if found == "" {
		return nil, nil
	}
	data, err := ioutil.ReadFile(found)
	if err != nil {
		return nil, err
	}
	return &file{baseDir, filename, data}, nil
}

func (f *file) Ast() (*ast.File, error) {
	loader := loader.Config{ParserMode: parser.ParseComments}
	astf, err := loader.ParseFile(f.filename, string(f.data))
	if err != nil {
		return nil, err
	}
	return astf, nil
}

func (f *file) GeneratedPath(factory string) string {
	return ""
}

func findpath(dir, target string) string {
	infolist, err := ioutil.ReadDir(dir)
	if err != nil {
		return ""
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
			return p
		}
	}
	for _, d := range subdirs {
		if found := findpath(d, target); found != "" {
			return found
		}
	}
	return ""
}
