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

var version = "dev"

const name = ".atlantis-conf.yaml"

func main() {
	targetPath := flag.String("path", ".", "PATH where to search for terragrunt.hcl files")
	showVersion := flag.Bool("version", false, "show version")
	flag.Parse()
	log.SetOutput(os.Stdout)

	if *showVersion {
		log.Println(version)
		os.Exit(0)
	}

	searchPath, err := common.NormalisePath(*targetPath)
	if err != nil {
		log.Fatal(err)
	}

	output := parser.ParseDir(searchPath)
	if len(output) == 0 {
		log.Fatalf("no terragrunt.hcl found in %s dir", *targetPath)
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
