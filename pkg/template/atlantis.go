package template

import (
	"fmt"
	"io"
	"text/template"
)

// Generate create `atlantis` config base on items.
func Generate(file io.Writer, items map[string][]string, showVersion bool) error {
	fileTemplate := template.Must(
		template.New("").Parse(`
projects:
{{- range $path, $dependencies := .Items }}
- autoplan:
  enabled: true
  when_modified:
  - '*.hcl'
  - '*.tf*'
{{- range $dependencies}}
  - {{ . }}
{{- end }}
  dir: {{ $path }}
  workflow: terragrunt
{{- end }}
{{- if .ShowVersion }}
version: 3
{{- end}}
`))

	if err := fileTemplate.Execute(file, struct {
		Items       map[string][]string
		ShowVersion bool
	}{
		Items:       items,
		ShowVersion: showVersion,
	}); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
