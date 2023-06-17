package template

import (
	"fmt"
	"io"
	"text/template"
)

// Generate create `atlantis` config base on items.
func Generate(file io.Writer, items map[string][]string, extraConfig string) error {
	fileTemplate := template.Must(
		template.New("").Parse(`
{{- if .Config }}
{{- .Config }}projects:
{{- else }}projects:
{{- end}}
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
`))

	if err := fileTemplate.Execute(file, struct {
		Config string
		Items  map[string][]string
	}{
		Config: extraConfig,
		Items:  items,
	}); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
