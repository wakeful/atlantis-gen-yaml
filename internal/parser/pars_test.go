// Copyright 2025 variHQ OÜ
// SPDX-License-Identifier: BSD-3-Clause

package parser

import (
	"reflect"
	"testing"
)

func TestParseDir(t *testing.T) {
	tests := []struct {
		name string
		path string
		want map[string][]string
	}{
		{
			name: "non existing dir should fail with empty map",
			path: "./non-existing",
			want: map[string][]string{},
		},
		{
			name: "files without dep. should return a map with empty string slice",
			path: "./empty-file",
			want: map[string][]string{
				"empty-file": nil,
			},
		},
		{
			path: "./",
			want: map[string][]string{
				"empty-file": nil,
				"with-dep": {
					"../../vpc",
					"../sg",
					"../src",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseDir(tt.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
