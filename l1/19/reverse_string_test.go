package main

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Single character",
			input:    "a",
			expected: "a",
		},
		{
			name:     "ASCII characters",
			input:    "hello",
			expected: "olleh",
		},
		{
			name:     "Unicode characters",
			input:    "привет",
			expected: "тевирп",
		},
		{
			name:     "Mixed Unicode",
			input:    "Hello, 世界!",
			expected: "!界世 ,olleH",
		},
		{
			name:     "Palindrome",
			input:    "madam",
			expected: "madam",
		},
		{
			name:     "Numbers",
			input:    "12345",
			expected: "54321",
		},
		{
			name:     "Special characters",
			input:    "!@#$%",
			expected: "%$#@!",
		},
		{
			name:     "Spaces",
			input:    "   ",
			expected: "   ",
		},
		{
			name:     "Combination",
			input:    "a1b2c3",
			expected: "3c2b1a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReverseString(tt.input)
			if result != tt.expected {
				t.Errorf("ReverseString(%q) = %q, want %q",
					tt.input, result, tt.expected)
			}
		})
	}
}
