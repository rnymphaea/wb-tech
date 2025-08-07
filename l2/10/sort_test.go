package main

import (
	"testing"
)

func TestCmpNumeric(t *testing.T) {
	tests := []struct {
		name string
		a, b line
		want int
	}{
		{
			name: "both numbers",
			a:    line{text: "a\t10", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "b\t5", key: 1, sep: "\t", reverse: 1},
			want: 1,
		},
		{
			name: "a missing column",
			a:    line{text: "apple", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "banana\t20", key: 1, sep: "\t", reverse: 1},
			want: -1,
		},
		{
			name: "b missing column",
			a:    line{text: "apple\t15", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "banana", key: 1, sep: "\t", reverse: 1},
			want: 1,
		},
		{
			name: "both missing columns",
			a:    line{text: "apple", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "banana", key: 1, sep: "\t", reverse: 1},
			want: -1,
		},
		{
			name: "a non-numeric",
			a:    line{text: "a\tfoo", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "b\t10", key: 1, sep: "\t", reverse: 1},
			want: -1,
		},
		{
			name: "b non-numeric",
			a:    line{text: "a\t15", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "b\tbar", key: 1, sep: "\t", reverse: 1},
			want: 1,
		},
		{
			name: "both non-numeric",
			a:    line{text: "a\tfoo", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "b\tbar", key: 1, sep: "\t", reverse: 1},
			want: -1,
		},
		{
			name: "equal numbers",
			a:    line{text: "a\t10", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "b\t10", key: 1, sep: "\t", reverse: 1},
			want: 0,
		},
		{
			name: "different separator",
			a:    line{text: "a,10", key: 1, sep: ",", reverse: 1},
			b:    line{text: "b,5", key: 1, sep: ",", reverse: 1},
			want: 1,
		},
		{
			name: "both numbers reverse",
			a:    line{text: "a\t10", key: 1, sep: "\t", reverse: -1},
			b:    line{text: "b\t5", key: 1, sep: "\t", reverse: -1},
			want: -1,
		},
		{
			name: "a missing column reverse",
			a:    line{text: "apple", key: 1, sep: "\t", reverse: -1},
			b:    line{text: "banana\t20", key: 1, sep: "\t", reverse: -1},
			want: 1,
		},
		{
			name: "b non-numeric reverse",
			a:    line{text: "a\t15", key: 1, sep: "\t", reverse: -1},
			b:    line{text: "b\tbar", key: 1, sep: "\t", reverse: -1},
			want: -1,
		},
		{
			name: "both non-numeric reverse",
			a:    line{text: "a\tfoo", key: 1, sep: "\t", reverse: -1},
			b:    line{text: "b\tbar", key: 1, sep: "\t", reverse: -1},
			want: 1,
		},
		{
			name: "equal numbers reverse",
			a:    line{text: "a\t10", key: 1, sep: "\t", reverse: -1},
			b:    line{text: "b\t10", key: 1, sep: "\t", reverse: -1},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cmpNumeric(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("cmpNumeric(%q, %q) = %d, want %d", tt.a.text, tt.b.text, got, tt.want)
			}
		})
	}
}
