// Copyright 2025 variHQ OÜ
// SPDX-License-Identifier: BSD-3-Clause

package common_test

import (
	"reflect"
	"testing"

	"github.com/wakeful/atlantis-gen-yaml/internal/common"
)

func TestFindFilesInPath(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			if got := common.FindFilesInPath(tt.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindFilesInPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
