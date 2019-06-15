package fopa

import (
	"io/ioutil"
	"path/filepath"
)

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
