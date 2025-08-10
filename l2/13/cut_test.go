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

func TestCut(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		opts  options
		want  []string
	}{
		{
			name:  "basic field selection",
			lines: []string{"a:b:c", "d:e:f"},
			opts: options{
				fields:          []int{1},
				inputDelimeter:  ":",
				outputDelimeter: ":",
				separatedOnly:   false,
			},
			want: []string{"b", "e"},
		},
		{
			name:  "multiple fields",
			lines: []string{"1,2,3,4", "a,b,c,d"},
			opts: options{
				fields:          []int{0, 2},
				inputDelimeter:  ",",
				outputDelimeter: "-",
				separatedOnly:   false,
			},
			want: []string{"1-3", "a-c"},
		},
		{
			name:  "separated only without delimiter",
			lines: []string{"no delimeter", "has:delimiter"},
			opts: options{
				fields:          []int{0},
				inputDelimeter:  ":",
				outputDelimeter: ":",
				separatedOnly:   true,
			},
			want: []string{"has"},
		},
		{
			name:  "separated only with delimiter",
			lines: []string{"skip:me", "keep:me"},
			opts: options{
				fields:          nil,
				inputDelimeter:  ":",
				outputDelimeter: ":",
				separatedOnly:   true,
			},
			want: []string{"skip:me", "keep:me"},
		},
		{
			name:  "no fields specified",
			lines: []string{"a:b:c", "d:e:f"},
			opts: options{
				fields:          nil,
				inputDelimeter:  ":",
				outputDelimeter: "-",
				separatedOnly:   false,
			},
			want: []string{"a-b-c", "d-e-f"},
		},
		{
			name:  "field out of range",
			lines: []string{"one:two", "three:four"},
			opts: options{
				fields:          []int{5},
				inputDelimeter:  ":",
				outputDelimeter: ":",
				separatedOnly:   false,
			},
			want: []string{"", ""},
		},
		{
			name:  "mixed existing and non-existing fields",
			lines: []string{"a:b:c", "d:e:f"},
			opts: options{
				fields:          []int{0, 5, 2},
				inputDelimeter:  ":",
				outputDelimeter: "-",
				separatedOnly:   false,
			},
			want: []string{"a-c", "d-f"},
		},
		{
			name:  "different input/output delimiters",
			lines: []string{"a,b,c", "d,e,f"},
			opts: options{
				fields:          []int{0, 2},
				inputDelimeter:  ",",
				outputDelimeter: "|",
				separatedOnly:   false,
			},
			want: []string{"a|c", "d|f"},
		},
		{
			name:  "single field with no delimiter",
			lines: []string{"no delimeter", "has,comma"},
			opts: options{
				fields:          []int{0},
				inputDelimeter:  ",",
				outputDelimeter: ",",
				separatedOnly:   false,
			},
			want: []string{"no delimeter", "has"},
		},
		{
			name:  "empty lines handling",
			lines: []string{"", "a:b", "", "c:d"},
			opts: options{
				fields:          []int{0},
				inputDelimeter:  ":",
				outputDelimeter: ":",
				separatedOnly:   false,
			},
			want: []string{"", "a", "", "c"},
		},
		{
			name:  "fields in reverse order",
			lines: []string{"a:b:c:d", "e:f:g:h"},
			opts: options{
				fields:          []int{3, 1},
				inputDelimeter:  ":",
				outputDelimeter: ":",
				separatedOnly:   false,
			},
			want: []string{"d:b", "h:f"},
		},
		{
			name:  "separated only with empty fields",
			lines: []string{":", "a:", ":b", "a:b"},
			opts: options{
				fields:          []int{0, 1},
				inputDelimeter:  ":",
				outputDelimeter: "-",
				separatedOnly:   true,
			},
			want: []string{"-", "a-", "-b", "a-b"},
		},
		{
			name:  "multiple character delimiters",
			lines: []string{"a||b||c", "d||e||f"},
			opts: options{
				fields:          []int{1},
				inputDelimeter:  "||",
				outputDelimeter: "--",
				separatedOnly:   false,
			},
			want: []string{"b", "e"},
		},
		{
			name:  "no fields with separated only",
			lines: []string{"nodlm", "has:dlm"},
			opts: options{
				fields:          []int{0},
				inputDelimeter:  ":",
				outputDelimeter: ":",
				separatedOnly:   true,
			},
			want: []string{"has"},
		},
		{
			name:  "all fields out of range",
			lines: []string{"a:b", "c:d"},
			opts: options{
				fields:          []int{5, 6},
				inputDelimeter:  ":",
				outputDelimeter: ":",
				separatedOnly:   false,
			},
			want: []string{"", ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cut(tt.lines, tt.opts)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cut() = %v, want %v", got, tt.want)
			}
		})
	}
}
