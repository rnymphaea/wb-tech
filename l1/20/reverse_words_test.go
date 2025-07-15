package main

import "testing"

func TestReverseWordsOrder(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "Empty slice",
			input:    []string{},
			expected: "",
		},
		{
			name:     "Single word",
			input:    []string{"hello"},
			expected: "hello",
		},
		{
			name:     "Two words",
			input:    []string{"hello", "world"},
			expected: "world hello",
		},
		{
			name:     "Three words",
			input:    []string{"this", "is", "test"},
			expected: "test is this",
		},
		{
			name:     "Multiple words",
			input:    []string{"go", "is", "an", "awesome", "language"},
			expected: "language awesome an is go",
		},
		{
			name:     "With empty strings",
			input:    []string{"", "hello", "", "world"},
			expected: "world  hello ",
		},
		{
			name:     "All empty strings",
			input:    []string{"", "", ""},
			expected: "  ",
		},
		{
			name:     "Mixed case words",
			input:    []string{"Hello", "WORLD", "gO"},
			expected: "gO WORLD Hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReverseWordsOrder(tt.input)
			if result != tt.expected {
				t.Errorf("ReverseWordsOrder(%v) = %q, want %q",
					tt.input, result, tt.expected)
			}
		})
	}
}
