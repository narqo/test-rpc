package main

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"path/filepath"
	"strings"
)

// Generator generates method to satisfy runtime.URLUnmarshaler interface.
type Generator struct {
	name       string
	targetName string
	namePkg    string
	pkg        *types.Package
	target     *types.Struct
}

func (g *Generator) parsePackageDir(dir string) error {
	pkg, err := build.Default.ImportDir(dir, 0)
	if err != nil {
		return fmt.Errorf("cannot process directory %s: %s", dir, err)
	}
	var names []string
	names = append(names, pkg.GoFiles...)
	names = append(names, pkg.CgoFiles...)
	names = prefixDirectory(dir, names)
	return g.parsePackage(dir, names, nil)
}

func (g *Generator) parsePackageFiles(names []string) {
	g.parsePackage(".", names, nil)
}

func prefixDirectory(dir string, names []string) []string {
	if dir == "." {
		return names
	}
	ret := make([]string, len(names))
	for i, name := range names {
		ret[i] = filepath.Join(dir, name)
	}
	return ret
}

type File struct {
	file *ast.File
	pkg *types.Package
}

func (g *Generator) parsePackage(dir string, names []string, text interface{}) error {
	var files []*File
	var astFiles []*ast.File
	g.pkg = new(types.Package)
	fs := token.NewFileSet()
	for _, name := range names {
		if !strings.HasSuffix(name, ".go") {
			continue
		}
		parsedFile, err := parser.ParseFile(fs, name, text, 0)
		if err != nil {
			return fmt.Errorf("parsing package: %s: %s", name, err)
		}
		astFiles = append(astFiles, parsedFile)
		files = append(files, &File{
			file: parsedFile,
			pkg:  g.pkg,
		})
	}
	if len(astFiles) == 0 {
		log.Fatalf("%s: no buildable Go files", dir)
	}
	log.Printf("ast files: %+v\n", *astFiles[0])
	//log.Printf("files: %+v\n", files)
	//g.pkg.name = astFiles[0].Name.Name
	//g.pkg.files = files
	//g.pkg.dir = dir
	//// Type check the package.
	//g.pkg.check(fs, astFiles)
	return nil
}

func (g *Generator) generate(typeName string) error {
	return nil
}
