package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	typeNames = flag.String("type", "", "list of type names")
)

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = Usage
	flag.Parse()
	if len(*typeNames) == 0 {
		flag.Usage()
		os.Exit(2)
	}
	types := strings.Split(*typeNames, ",")

	args := flag.Args()
	if len(args) == 0 {
		args = []string{"."}
	}

	var g Generator
	if len(args) == 1 && isDirectory(args[0]) {
		g.parsePackageDir(args[0])
	} else {
		g.parsePackageFiles(args)
	}

	for _, typeName := range types {
		g.generate(typeName)
	}
}

func isDirectory(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}
	return info.IsDir()
}
