// Copyright 2025 variHQ OÜ
// SPDX-License-Identifier: BSD-3-Clause

package terragrunt

import (
	"context"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/gruntwork-io/terragrunt/config"
	"github.com/gruntwork-io/terragrunt/options"
	"github.com/gruntwork-io/terragrunt/pkg/log"
	"github.com/gruntwork-io/terragrunt/pkg/log/format"
)

// GetDependencies parse given file and return the PATH(s) to module(s) it depends on.
func GetDependencies(path string) ([]string, error) {
	terragruntOptions, err := options.NewTerragruntOptionsWithConfigPath(path)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	nullLogger := log.New(
		log.WithOutput(io.Discard),
		log.WithLevel(log.InfoLevel),
		log.WithFormatter(format.NewFormatter(format.NewKeyValueFormatPlaceholders())),
	)

	ctx := config.NewParsingContext(context.Background(), nullLogger, terragruntOptions)

	configFile, err := config.PartialParseConfigFile(ctx.WithDecodeList(
		config.DependencyBlock,
		config.DependenciesBlock,
		config.TerraformSource,
	), nullLogger, path, nil)
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
