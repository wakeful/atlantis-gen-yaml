// Copyright 2025 variHQ OÜ
// SPDX-License-Identifier: BSD-3-Clause

package common

import "testing"

func TestNormalisePath(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		want    string
		wantErr bool
	}{
		{
			path:    "./tmp",
			want:    "./tmp/",
			wantErr: false,
		},
		{
			path:    "/tmp",
			want:    "../../../../../../../tmp/",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NormalisePath(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("NormalisePath() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if got != tt.want {
				t.Errorf("NormalisePath() got = %v, want %v", got, tt.want)
			}
		})
	}
}
