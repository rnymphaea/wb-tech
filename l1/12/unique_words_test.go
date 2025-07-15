package main

import (
	"reflect"
	"testing"
)

func TestGetUniqueWords(t *testing.T) {
	tests := []struct {
		name       string
		input      []string
		ignoreCase bool
		expected   []string
	}{
		{
			name:       "Case sensitive",
			input:      []string{"a", "A", "b", "a"},
			ignoreCase: false,
			expected:   []string{"a", "A", "b"},
		},
		{
			name:       "Ignore case",
			input:      []string{"a", "A", "b", "a"},
			ignoreCase: true,
			expected:   []string{"a", "b"},
		},
		{
			name:       "Empty slice",
			input:      []string{},
			ignoreCase: false,
			expected:   []string{},
		},
		{
			name:       "All duplicates",
			input:      []string{"x", "x", "x"},
			ignoreCase: false,
			expected:   []string{"x"},
		},
		{
			name:       "Mixed words",
			input:      []string{"Go", "go", "GO", "Python", "python"},
			ignoreCase: true,
			expected:   []string{"go", "python"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetUniqueWords(tt.input, tt.ignoreCase)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("GetUniqueWords() = %v, want %v", result, tt.expected)
			}
		})
	}
}
