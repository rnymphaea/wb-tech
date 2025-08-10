package main

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestGetFields(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []int
		wantErr bool
	}{
		{
			name:    "single field",
			input:   "5",
			want:    []int{4},
			wantErr: false,
		},
		{
			name:    "multiple fields",
			input:   "1,3,5",
			want:    []int{0, 2, 4},
			wantErr: false,
		},
		{
			name:    "simple range",
			input:   "3-5",
			want:    []int{2, 3, 4},
			wantErr: false,
		},
		{
			name:    "reverse range",
			input:   "5-3",
			want:    []int{2, 3, 4},
			wantErr: false,
		},
		{
			name:    "mixed fields and ranges",
			input:   "1,3-5,7",
			want:    []int{0, 2, 3, 4, 6},
			wantErr: false,
		},
		{
			name:    "single field zero",
			input:   "0",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "negative field",
			input:   "-2",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid range format",
			input:   "1-3-5",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "non-numeric field",
			input:   "a",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "non-numeric in range start",
			input:   "a-5",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "non-numeric in range end",
			input:   "3-b",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty input",
			input:   "",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "trailing comma",
			input:   "1,2,",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "double dash",
			input:   "1--3",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "range with negative end",
			input:   "3--5",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "large range",
			input:   "1-5",
			want:    []int{0, 1, 2, 3, 4},
			wantErr: false,
		},
		{
			name:    "single field at max int",
			input:   fmt.Sprintf("%d", math.MaxInt),
			want:    []int{math.MaxInt - 1},
			wantErr: false,
		},
		{
			name:    "range crossing max int",
			input:   fmt.Sprintf("%d-%d", math.MaxInt-2, math.MaxInt),
			want:    []int{math.MaxInt - 3, math.MaxInt - 2, math.MaxInt - 1},
			wantErr: false,
		},
		{
			name:    "comma only",
			input:   ",",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "dash only",
			input:   "-",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getFields(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("getFields() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFields() = %v, want %v", got, tt.want)
			}
		})
	}
}
