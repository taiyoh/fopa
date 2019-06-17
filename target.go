package fopa

// Target represents code generation target.
type Target struct {
	file *File
	ast  *syntaxTree
}

// FindTarget provides Target object from supplied filename and type definition.
func FindTarget(base, basedir, filename string) (*Target, error) {
	file, err := findFile(basedir, filename)
	if file == nil {
		return nil, err
	}
	astf, err := file.AST()
	if err != nil {
		return nil, err
	}

	synTree, err := findAst(base, astf)
	if err != nil {
		return nil, err
	}
	return &Target{
		file: file,
		ast:  synTree,
	}, nil
}

// Build returns generated go code.
func (t *Target) Build(pkgname, factory, builder string) *File {
	return buildFile(t.ast, t.file.baseDir, pkgname, factory, builder)
}
