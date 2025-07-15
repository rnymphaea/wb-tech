package main

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		reverse  bool
		expected []int
	}{
		{
			name:     "Empty array",
			input:    []int{},
			reverse:  false,
			expected: []int{},
		},
		{
			name:     "Single element",
			input:    []int{42},
			reverse:  false,
			expected: []int{42},
		},
		{
			name:     "Already sorted ascending",
			input:    []int{1, 2, 3, 4, 5},
			reverse:  false,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Unsorted array",
			input:    []int{5, 3, 1, 4, 2},
			reverse:  false,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted",
			input:    []int{5, 4, 3, 2, 1},
			reverse:  false,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "With duplicates",
			input:    []int{3, 1, 2, 3, 1},
			reverse:  false,
			expected: []int{1, 1, 2, 3, 3},
		},
		{
			name:     "Reverse order true",
			input:    []int{1, 3, 2, 5, 4},
			reverse:  true,
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "All equal elements",
			input:    []int{7, 7, 7, 7},
			reverse:  false,
			expected: []int{7, 7, 7, 7},
		},
		{
			name:     "Large random array",
			input:    []int{9, 4, 7, 2, 5, 8, 1, 6, 3, 0},
			reverse:  false,
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := QuickSort(tt.input, tt.reverse)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("QuickSort(%v, %t) = %v, want %v",
					tt.input, tt.reverse, result, tt.expected)
			}
		})
	}
}
