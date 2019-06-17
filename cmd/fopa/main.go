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
	f := target.Build(v.pkgName, v.factory, v.builder)
	if v.stdout {
		fmt.Print(string(f.Raw()))
		os.Exit(0)
	}
	ioutil.WriteFile(f.FilePath(), f.Raw(), 0644)
}

type vars struct {
	base     string
	factory  string
	builder  string
	stdout   bool
	filename string
	pkgName  string
	pwd      string
}

func initialize() (*vars, error) {
	base := ""
	factory := ""
	builder := ""
	stdout := false
	flag.StringVar(&base, "base", "", "base struct for target")
	flag.StringVar(&factory, "factory", "", "fill factory struct name if not {.base}Factory")
	flag.StringVar(&builder, "builder", "", "build struct if already defined")
	flag.BoolVar(&stdout, "stdout", false, "set true if debug")
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
		stdout:   stdout,
		filename: filename,
		pkgName:  pkgName,
		pwd:      pwd,
	}
	return v, nil
}
