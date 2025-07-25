package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "simple unpack",
			input:   "a4bc2d5e",
			want:    "aaaabccddddde",
			wantErr: false,
		},
		{
			name:    "no digits",
			input:   "abcd",
			want:    "abcd",
			wantErr: false,
		},
		{
			name:    "empty string",
			input:   "",
			want:    "",
			wantErr: false,
		},
		{
			name:    "only digits",
			input:   "45",
			want:    "",
			wantErr: true,
		},
		{
			name:    "starts with digit",
			input:   "4a",
			want:    "",
			wantErr: true,
		},
		{
			name:    "multiple digits in a row",
			input:   "a10b",
			want:    "aaaaaaaaaab",
			wantErr: false,
		},
		{
			name:    "unicode characters",
			input:   "ф2ы3",
			want:    "ффыыы",
			wantErr: false,
		},
		{
			name:    "single character",
			input:   "a",
			want:    "a",
			wantErr: false,
		},
		{
			name:    "single digit after letter",
			input:   "a1",
			want:    "a",
			wantErr: false,
		},
		{
			name:    "digit zero",
			input:   "a0",
			want:    "",
			wantErr: false,
		},
		{
			name:    "multiple zeros",
			input:   "a0b0c0",
			want:    "",
			wantErr: false,
		},
		{
			name:    "mixed letters and digits",
			input:   "a2b3c1d0e1",
			want:    "aabbbce",
			wantErr: false,
		},
		{
			name:    "digits at end",
			input:   "abc10",
			want:    "abcccccccccc",
			wantErr: false,
		},
		{
			name:    "digits in middle",
			input:   "a2b3c",
			want:    "aabbbc",
			wantErr: false,
		},
		{
			name:    "escaped digits",
			input:   "qwe\\4\\5",
			want:    "qwe45",
			wantErr: false,
		},
		{
			name:    "mixed escaped and unescaped digits",
			input:   "qwe\\45",
			want:    "qwe44444",
			wantErr: false,
		},
		{
			name:    "complex escape sequence",
			input:   "a\\2\\3b4\\5",
			want:    "a23bbbb5",
			wantErr: false,
		},
		{
			name:    "escape at start",
			input:   "\\2a3",
			want:    "2aaa",
			wantErr: false,
		},
		{
			name:    "escape at end",
			input:   "a3\\2",
			want:    "aaa2",
			wantErr: false,
		},
		{
			name:    "invalid escape sequence",
			input:   "a\\",
			want:    "a\\",
			wantErr: false,
		},
		{
			name:    "escape before letter",
			input:   "a\\bc3",
			want:    "a\\bccc",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unpack(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("unpack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("str = %s, unpack() = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
