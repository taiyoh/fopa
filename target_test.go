package fopa_test

import (
	"os"
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
}
