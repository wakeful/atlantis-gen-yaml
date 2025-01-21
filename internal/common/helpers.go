// Copyright 2025 variHQ OÜ
// SPDX-License-Identifier: BSD-3-Clause

package common

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// NormalisePath convert given path to relative version.
func NormalisePath(path string) (string, error) {
	if !filepath.IsAbs(path) {
		return withSuffix(path), nil
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	path, err = filepath.Rel(currentDir, path)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	return withSuffix(path), nil
}

func withSuffix(path string) string {
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	return path
}
