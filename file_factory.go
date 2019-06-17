package fopa

import (
	"io/ioutil"
	"path/filepath"
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
