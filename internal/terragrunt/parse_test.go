// Copyright 2025 variHQ OÜ
// SPDX-License-Identifier: BSD-3-Clause

package terragrunt

import (
	"reflect"
	"testing"
)

func TestGetDependencies(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		want    []string
		wantErr bool
	}{
		{
			name:    "no dir should give back an error",
			wantErr: true,
		},
		{
			name:    "empty file should return an empty map",
			path:    "./empty-file/terragrunt.hcl",
			wantErr: false,
			want:    nil,
		},
		{
			name:    "invalid syntax should error",
			path:    "./invalid-syntax/terragrunt.hcl",
			wantErr: true,
		},
		{
			name:    "we should get a map of string slices",
			path:    "with-dep/terragrunt.hcl",
			wantErr: false,
			want:    []string{"../../vpc", "../sg", "../src"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDependencies(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDependencies() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDependencies() got = %v, want %v", got, tt.want)
			}
		})
	}
}
