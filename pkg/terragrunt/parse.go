package terragrunt

import (
	"fmt"
	"sort"
	"strings"

	"github.com/gruntwork-io/terragrunt/config"
	"github.com/gruntwork-io/terragrunt/options"
)

// GetDependencies parse given file and return the PATH(s) to module(s) it depends on.
func GetDependencies(path string) ([]string, error) {
	terragruntOptions, err := options.NewTerragruntOptionsWithConfigPath(path)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	configFile, err := config.PartialParseConfigFile(path, terragruntOptions, nil, []config.PartialDecodeSectionType{
		config.DependencyBlock,
		config.DependenciesBlock,
		config.TerraformSource,
	})
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	var output []string

	if configFile.Terraform != nil {
		source := configFile.Terraform.Source
		if source != nil && *source != "" {
			if strings.HasPrefix(*source, "./") || strings.HasPrefix(*source, "../") {
				output = append(output, *source)
			}
		}
	}

	if configFile.Dependencies != nil {
		output = append(output, configFile.Dependencies.Paths...)
	}

	sort.Strings(output)

	return output, nil
}
