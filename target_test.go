package fopa_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/taiyoh/fopa"
)

func TestTargetFail(t *testing.T) {
	pwd, _ := os.Getwd()
	for idx, tt := range []struct {
		typ    string
		name   string
		exists bool
	}{
		{"test1", "dummi_test.go", false},
		{"test2", "dummy_test.go", true},
	} {
		tgt, err := fopa.FindTarget(tt.typ, pwd, tt.name)
		if tgt != nil {
			t.Errorf("[%d] target found", idx)
		}
		if tt.exists && err == nil {
			t.Errorf("[%d] error not found", idx)
		}
	}
}

func TestTargetSuccess(t *testing.T) {
	pwd, _ := os.Getwd()
	tgt, err := fopa.FindTarget("test1", pwd, "dummy_test.go")
	if tgt == nil {
		t.Errorf("target not found")
	}
	if err != nil {
		t.Errorf("error found")
	}
	output := tgt.Build("pkg1", "test1Factory", "test1BuilderFn")
	ref := `// Code generated by fopa. DO NOT EDIT.

package pkg1

type test1BuilderFn func(*test1)

func (f *test1Factory) SetupTest1(fns ...test1BuilderFn) *test1 {
	o := &test1{}
	for _, fn := range fns {
		fn(o)
	}
	return o
}

func (f *test1Factory) FillBbb(v int) test1BuilderFn {
	return func(p *test1) {
		p.bbb = tt(v)
	}
}

func (f *test1Factory) FillCcc(v string) test1BuilderFn {
	return func(p *test1) {
		p.ccc = hoge{tt2(v)}
	}
}

func (f *test1Factory) FillDdd(v int) test1BuilderFn {
	return func(p *test1) {
		p.ddd = v
	}
}
`
	if output != ref {
		t.Error("generated code is something wrong")
	}

	gen := tgt.GeneratedPath()
	if gen != filepath.Join(pwd, "internal", "pkg1", "dummy_test_gen.go") {
		t.Errorf("wrong generated path returns: %s", gen)
	}
}
