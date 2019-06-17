package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/taiyoh/fopa"
)

func main() {
	v, err := initialize()
	if err != nil {
		panic(err)
	}
	target, err := fopa.FindTarget(v.base, v.pwd, v.filename)
	if err != nil {
		panic(err)
	}
	output := target.Build(v.pkgName, v.factory, v.builder)
	ioutil.WriteFile(target.GeneratedPath(v.factory), output, 0644)
}

type vars struct {
	base     string
	factory  string
	builder  string
	filename string
	pkgName  string
	pwd      string
}

func initialize() (*vars, error) {
	base := ""
	factory := ""
	builder := ""
	flag.StringVar(&base, "base", "", "base struct")
	flag.StringVar(&factory, "factory", "", "factory struct")
	flag.StringVar(&builder, "builder", "", "factory struct")
	flag.Parse()

	if base == "" {
		return nil, errors.New("base not found")
	}
	if factory == "" {
		factory = fmt.Sprintf("%sFactory", base)
	}
	if builder == "" {
		builder = fmt.Sprintf("%sBuilderFn", base)
	}
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	filename := os.Getenv("GOFILE")
	if filename == "" {
		return nil, errors.New("GOFILE requires")
	}
	pkgName := os.Getenv("GOPACKAGE")
	if pkgName == "" {
		return nil, errors.New("GOPACKAGE requires")
	}
	v := &vars{
		base:     base,
		factory:  factory,
		builder:  builder,
		filename: filename,
		pkgName:  pkgName,
		pwd:      pwd,
	}
	return v, nil
}
