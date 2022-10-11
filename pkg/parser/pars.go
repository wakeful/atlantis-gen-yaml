package parser

import (
	"log"
	"strings"

	"github.com/wakeful/atlantis-gen-yaml/pkg/common"
	"github.com/wakeful/atlantis-gen-yaml/pkg/terragrunt"
)

// ParseDir locate all `terragrunt.hcl` file(s) in given workDir and return them with list of dependencies.
func ParseDir(workDir string) map[string][]string {
	filesInPath := common.FindFilesInPath(workDir)
	output := map[string][]string{}

	for _, file := range filesInPath {
		dependencies, err := terragrunt.GetDependencies(file)
		if err != nil {
			log.Println(err)
		}

		dest := strings.TrimSuffix(file, "/terragrunt.hcl")

		output[dest] = dependencies
	}

	return output
}
