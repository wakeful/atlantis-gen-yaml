package common

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// FindFilesInPath locate all terragrunt.hcl files in given PATH recursively.
func FindFilesInPath(root string) []string {
	var files []string

	_ = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)

			return nil
		}
		if !info.IsDir() && info.Name() == "terragrunt.hcl" {
			if strings.Contains(path, "/.terragrunt-cache/") {
				return nil
			}
			files = append(files, path)
		}

		return nil
	})

	return files
}
