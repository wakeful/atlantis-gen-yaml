package template

import (
	"bytes"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name        string
		items       map[string][]string
		wantFile    string
		wantErr     bool
		showVersion bool
	}{
		{
			wantErr:     false,
			showVersion: true,
			items: map[string][]string{
				"empty-file": nil,
			},
			wantFile: `
projects:
- autoplan:
  enabled: true
  when_modified:
  - '*.hcl'
  - '*.tf*'
  dir: empty-file
  workflow: terragrunt
version: 3
`,
		},
		{
			wantErr:     false,
			showVersion: false,
			items: map[string][]string{
				"empty-file": nil,
				"with-dep": {
					"../../vpc",
					"../sg",
				},
			},
			wantFile: `
projects:
- autoplan:
  enabled: true
  when_modified:
  - '*.hcl'
  - '*.tf*'
  dir: empty-file
  workflow: terragrunt
- autoplan:
  enabled: true
  when_modified:
  - '*.hcl'
  - '*.tf*'
  - ../../vpc
  - ../sg
  dir: with-dep
  workflow: terragrunt
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := &bytes.Buffer{}
			err := Generate(file, tt.items, tt.showVersion)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if gotFile := file.String(); gotFile != tt.wantFile {
				t.Errorf("Generate() gotFile = %v, want %v", gotFile, tt.wantFile)
			}
		})
	}
}
