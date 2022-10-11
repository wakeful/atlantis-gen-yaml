package common

import (
	"reflect"
	"testing"
)

func TestFindFilesInPath(t *testing.T) {
	tests := []struct {
		name string
		path string
		want []string
	}{
		{
			name: "empty dir",
			path: "./empty",
			want: nil,
		},
		{
			name: "recursive check",
			path: "./rec",
			want: []string{
				"rec/com01/terragrunt.hcl",
				"rec/com02/terragrunt.hcl",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindFilesInPath(tt.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindFilesInPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
