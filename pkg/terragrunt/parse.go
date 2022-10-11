package terragrunt

import (
	"fmt"
	"sort"

	"github.com/gruntwork-io/terragrunt/config"
	"github.com/gruntwork-io/terragrunt/options"
)

// GetDependencies parse given file and return the PATH(s) to module(s) it depends on.
func GetDependencies(path string) ([]string, error) {
	terragruntOptions, err := options.NewTerragruntOptions(path)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	configFile, err := config.PartialParseConfigFile(path, terragruntOptions, nil, []config.PartialDecodeSectionType{
		config.DependencyBlock,
		config.DependenciesBlock,
	})
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	if configFile.Dependencies == nil {
		return nil, nil
	}

	output := configFile.Dependencies.Paths
	sort.Strings(output)

	return output, nil
}
