// Copyright 2025 variHQ OÜ
// SPDX-License-Identifier: BSD-3-Clause

package parser

import (
	"log"
	"strings"

	"github.com/wakeful/atlantis-gen-yaml/internal/common"
	"github.com/wakeful/atlantis-gen-yaml/internal/terragrunt"
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
