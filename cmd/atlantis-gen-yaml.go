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
	targetPath  string
	showVersion bool
	version     = "dev"
)

func init() {
	flag.StringVar(&targetPath, "path", ".", "PATH where to search for terragrunt.hcl files")
	flag.BoolVar(&showVersion, "version", false, "show version")
}

const name = ".atlantis-conf.yaml"

func main() {
	flag.Parse()

	if showVersion {
		log.Println(version)
		os.Exit(0)
	}

	searchPath, err := common.NormalisePath(targetPath)
	if err != nil {
		log.Fatal(err)
	}

	output := parser.ParseDir(searchPath)
	if len(output) == 0 {
		log.Fatalf("no terragrunt.hcl found in %s dir", targetPath)
	}

	var extraConfig string

	if _, errReadConfFile := os.Stat(name); errReadConfFile == nil {
		buff, _ := os.ReadFile(name)
		extraConfig = string(buff)
	}

	if err := template.Generate(os.Stdout, output, extraConfig); err != nil {
		log.Fatal(fmt.Errorf("%w", err))
	}
}
