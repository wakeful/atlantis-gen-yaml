package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/wakeful/atlantis-gen-yaml/pkg/common"
	"github.com/wakeful/atlantis-gen-yaml/pkg/parser"
	"github.com/wakeful/atlantis-gen-yaml/pkg/template"
)

var (
	showVersion bool
	targetPath  string
)

func init() {
	flag.StringVar(&targetPath, "path", ".", "PATH where to search for terragrunt.hcl files")
	flag.BoolVar(&showVersion, "show-version", false, "should we include atlantis version in output?")
}

func main() {
	flag.Parse()

	searchPath, err := common.NormalisePath(targetPath)
	if err != nil {
		log.Fatal(err)
	}

	output := parser.ParseDir(searchPath)
	if len(output) == 0 {
		log.Fatalf("no terragrunt.hcl found in %s dir", targetPath)
	}

	if err := template.Generate(os.Stdout, output, showVersion); err != nil {
		log.Fatal(fmt.Errorf("%w", err))
	}
}
