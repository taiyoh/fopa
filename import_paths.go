package fopa

import "fmt"

type importPath struct {
	sig      string
	exported bool
	order    int
	path     string
}

func (i importPath) string() string {
	if i.exported {
		return fmt.Sprintf("%s \"%s\"", i.sig, i.path)
	}
	return fmt.Sprintf(`"%s"`, i.path)
}
