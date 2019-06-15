package fopa

import "strings"

type tag struct {
	raw        string
	enabled    bool
	acceptType string
	fillExpr   string
}

func newTag(raw string) tag {
	t := &tag{raw: raw, enabled: true}
	parts := strings.Split(raw, ";")
	for _, p := range parts {
		if p == "false" {
			return tag{raw: raw, enabled: false}
		}
		kv := strings.Split(p, ":")
		if len(kv) != 2 {
			continue
		}
		switch kv[0] {
		case "accept":
			t.acceptType = kv[1]
		case "expr":
			t.fillExpr = kv[1]
		}
	}
	return *t
}

func (t tag) acceptArgCount() int {
	return len(strings.Split(t.acceptType, ","))
}

func (t tag) exprStepsCount() int {
	return len(strings.Split(t.fillExpr, ";"))
}
