package main

import (
	"testing"
)

func TestBinarySearchBasic(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		target int
		want   int
	}{
		{"Element exists", []int{1, 3, 5, 7, 9}, 5, 2},
		{"First element", []int{1, 3, 5, 7, 9}, 1, 0},
		{"Last element", []int{1, 3, 5, 7, 9}, 9, 4},
		{"Element not exists", []int{1, 3, 5, 7, 9}, 4, -1},
		{"Empty array", []int{}, 1, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.arr, tt.target); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchGeneric(t *testing.T) {
	t.Run("String slice", func(t *testing.T) {
		arr := []string{"apple", "banana", "orange"}
		if got := BinarySearch(arr, "banana"); got != 1 {
			t.Errorf("BinarySearch(%v, 'banana') = %v, want 1", arr, got)
		}
	})

	t.Run("Float slice", func(t *testing.T) {
		arr := []float64{1.1, 2.2, 3.3}
		if got := BinarySearch(arr, 2.2); got != 1 {
			t.Errorf("BinarySearch(%v, %v) = %v, want 1", arr, 2.2, got)
		}
	})
}

func TestBinarySearchEdgeCases(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		target int
		want   int
	}{
		{"Single element exists", []int{42}, 42, 0},
		{"Single element not exists", []int{42}, 43, -1},
		{"Large array", makeRange(1, 1000000), 999999, 999998},
		{"All same elements", []int{5, 5, 5, 5}, 5, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.arr, tt.target); got != tt.want {
				t.Errorf("BinarySearch(%v, %v) = %v, want %v", tt.arr, tt.target, got, tt.want)
			}
		})
	}
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
