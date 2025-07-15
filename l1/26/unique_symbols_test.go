package main

import "testing"

func TestUniqueSymbols(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: true,
		},
		{
			name:     "All unique",
			input:    "abcdef",
			expected: true,
		},
		{
			name:     "All duplicates",
			input:    "aaaaa",
			expected: false,
		},
		{
			name:     "Mixed case unique",
			input:    "AbCdEf",
			expected: true,
		},
		{
			name:     "Mixed case duplicates",
			input:    "AbCdefA",
			expected: false,
		},
		{
			name:     "Special characters",
			input:    "!@#$%",
			expected: true,
		},
		{
			name:     "Unicode characters",
			input:    "привет",
			expected: true,
		},
		{
			name:     "Unicode duplicates",
			input:    "приветп",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := UniqueSymbols(tt.input)
			if result != tt.expected {
				t.Errorf("UniqueSymbols(%q) = %v, want %v",
					tt.input, result, tt.expected)
			}
		})
	}
}
