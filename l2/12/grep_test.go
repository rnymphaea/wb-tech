package main

import (
	"reflect"
	"regexp"
	"testing"
)

func TestGrepRE(t *testing.T) {
	tests := []struct {
		name   string
		lines  []string
		re     *regexp.Regexp
		invert bool
		want   []bool
	}{
		{
			name:   "basic match",
			lines:  []string{"hello", "world", "golang"},
			re:     regexp.MustCompile("go"),
			invert: false,
			want:   []bool{false, false, true},
		},
		{
			name:   "case sensitive",
			lines:  []string{"Go", "go", "GO"},
			re:     regexp.MustCompile("go"),
			invert: false,
			want:   []bool{false, true, false},
		},
		{
			name:   "invert match",
			lines:  []string{"cat", "dog", "fish"},
			re:     regexp.MustCompile("dog"),
			invert: true,
			want:   []bool{true, false, true},
		},
		{
			name:   "empty lines",
			lines:  []string{"", " ", "\t"},
			re:     regexp.MustCompile("^$"),
			invert: false,
			want:   []bool{true, false, false},
		},
		{
			name:   "special characters",
			lines:  []string{"a.b", "a+b", "a*b"},
			re:     regexp.MustCompile(`a\.b`),
			invert: false,
			want:   []bool{true, false, false},
		},
		{
			name:   "multiline match",
			lines:  []string{"start", "middle", "end"},
			re:     regexp.MustCompile("mid"),
			invert: false,
			want:   []bool{false, true, false},
		},
		{
			name:   "whole word",
			lines:  []string{"go", "golang", "ago"},
			re:     regexp.MustCompile(`\bgo\b`),
			invert: false,
			want:   []bool{true, false, false},
		},
		{
			name:   "case insensitive flag",
			lines:  []string{"Go", "go", "GO"},
			re:     regexp.MustCompile(`(?i)go`),
			invert: false,
			want:   []bool{true, true, true},
		},
		{
			name:   "invert with multiple matches",
			lines:  []string{"a", "b", "a", "c"},
			re:     regexp.MustCompile("a"),
			invert: true,
			want:   []bool{false, true, false, true},
		},
		{
			name:   "empty pattern",
			lines:  []string{"any", "string"},
			re:     regexp.MustCompile(""),
			invert: false,
			want:   []bool{true, true},
		},
		{
			name:   "start anchor",
			lines:  []string{"start", "not start", "start again"},
			re:     regexp.MustCompile(`^start`),
			invert: false,
			want:   []bool{true, false, true},
		},
		{
			name:   "complex regex",
			lines:  []string{"abc123", "def456", "xyz789"},
			re:     regexp.MustCompile(`[a-z]+\d{3}`),
			invert: false,
			want:   []bool{true, true, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := grepRE(tt.lines, tt.re, tt.invert)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("grepRE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrepStrict(t *testing.T) {
	tests := []struct {
		name       string
		lines      []string
		pattern    string
		ignoreCase bool
		invert     bool
		want       []bool
	}{
		{
			name:       "basic substring",
			lines:      []string{"hello world", "world peace", "peace"},
			pattern:    "world",
			ignoreCase: false,
			invert:     false,
			want:       []bool{true, true, false},
		},
		{
			name:       "ignore case",
			lines:      []string{"Go", "go", "GO"},
			pattern:    "go",
			ignoreCase: true,
			invert:     false,
			want:       []bool{true, true, true},
		},
		{
			name:       "case sensitive",
			lines:      []string{"Go", "go", "GO"},
			pattern:    "go",
			ignoreCase: false,
			invert:     false,
			want:       []bool{false, true, false},
		},
		{
			name:       "invert match",
			lines:      []string{"apple", "banana", "orange"},
			pattern:    "an",
			ignoreCase: false,
			invert:     true,
			want:       []bool{true, false, false},
		},
		{
			name:       "empty pattern",
			lines:      []string{"", "a", " "},
			pattern:    "",
			ignoreCase: false,
			invert:     false,
			want:       []bool{true, true, true},
		},
		{
			name:       "special characters",
			lines:      []string{"a.b", "a+b", "a*b"},
			pattern:    "a.b",
			ignoreCase: false,
			invert:     false,
			want:       []bool{true, false, false},
		},
		{
			name:       "whole word not match",
			lines:      []string{"go", "golang", "ago"},
			pattern:    "go",
			ignoreCase: false,
			invert:     false,
			want:       []bool{true, true, true},
		},
		{
			name:       "invert with ignore case",
			lines:      []string{"Go", "go", "GO"},
			pattern:    "go",
			ignoreCase: true,
			invert:     true,
			want:       []bool{false, false, false},
		},
		{
			name:       "pattern longer than string",
			lines:      []string{"short", "longer", "very long"},
			pattern:    "this is very long pattern",
			ignoreCase: false,
			invert:     false,
			want:       []bool{false, false, false},
		},
		{
			name:       "exact match required",
			lines:      []string{"exact", "exact match", "match exact"},
			pattern:    "exact",
			ignoreCase: false,
			invert:     false,
			want:       []bool{true, true, true},
		},
		{
			name:       "whitespace sensitivity",
			lines:      []string{" space ", "space", " sp ace "},
			pattern:    "space",
			ignoreCase: false,
			invert:     false,
			want:       []bool{true, true, false},
		},
		{
			name:       "unicode characters",
			lines:      []string{"日本語", "日本", "日本語です"},
			pattern:    "日本",
			ignoreCase: false,
			invert:     false,
			want:       []bool{true, true, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := grepStrict(tt.lines, tt.pattern, tt.ignoreCase, tt.invert)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("grepStrict() = %v, want %v", got, tt.want)
			}
		})
	}
}
