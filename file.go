package fopa

import (
	"go/ast"
	"go/parser"
	"path/filepath"

	"golang.org/x/tools/go/loader"
)

// File represents file abstraction in this tool.
type File struct {
	baseDir  string
	filename string
	data     []byte
}

// AST returns abstract syntax tree object.
func (f *File) AST() (*ast.File, error) {
	loader := loader.Config{ParserMode: parser.ParseComments}
	astf, err := loader.ParseFile(f.filename, string(f.data))
	if err != nil {
		return nil, err
	}
	return astf, nil
}

// FilePath returns absolute filepath string.
func (f *File) FilePath() string {
	return filepath.Join(f.baseDir, f.filename)
}

// Raw returns byte slice file data.
func (f *File) Raw() []byte {
	return f.data
}
