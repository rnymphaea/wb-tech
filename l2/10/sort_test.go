package main

import (
	"strings"
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

func TestCmpMonth(t *testing.T) {
	tests := []struct {
		name string
		a, b line
		want int
	}{
		{
			name: "same month, different days",
			a:    line{text: "event1\t15 Jan", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "event2\t10 Jan", key: 1, sep: "\t", reverse: 1},
			want: 5,
		},
		{
			name: "different months",
			a:    line{text: "event1\t10 Feb", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "event2\t15 Jan", key: 1, sep: "\t", reverse: 1},
			want: 1,
		},
		{
			name: "same date",
			a:    line{text: "event1\t10 Mar", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "event2\t10 Mar", key: 1, sep: "\t", reverse: 1},
			want: 0,
		},
		{
			name: "a invalid month",
			a:    line{text: "event1\t15 Inv", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "event2\t10 Jan", key: 1, sep: "\t", reverse: 1},
			want: -1,
		},
		{
			name: "b invalid month",
			a:    line{text: "event1\t15 Jan", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "event2\t10 Xyz", key: 1, sep: "\t", reverse: 1},
			want: 1,
		},
		{
			name: "both invalid months",
			a:    line{text: "event1\t15 Foo", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "event2\t10 Bar", key: 1, sep: "\t", reverse: 1},
			want: strings.Compare("event1\t15 Foo", "event2\t10 Bar"),
		},
		{
			name: "a invalid day format",
			a:    line{text: "event1\t1x5 Jan", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "event2\t10 Jan", key: 1, sep: "\t", reverse: 1},
			want: -1,
		},

		{
			name: "a missing column",
			a:    line{text: "event1", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "event2\t10 Jan", key: 1, sep: "\t", reverse: 1},
			want: -1,
		},
		{
			name: "b missing column",
			a:    line{text: "event1\t10 Jan", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "event2", key: 1, sep: "\t", reverse: 1},
			want: 1,
		},
		{
			name: "both missing columns",
			a:    line{text: "event1", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "event2", key: 1, sep: "\t", reverse: 1},
			want: strings.Compare("event1", "event2"),
		},

		{
			name: "different date formats",
			a:    line{text: "event1\t25DEC", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "event2\t10 Dec", key: 1, sep: "\t", reverse: 1},
			want: strings.Compare("event1", "event2"),
		},
		{
			name: "single word instead of date",
			a:    line{text: "event1\tJanuary", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "event2\tFebruary", key: 1, sep: "\t", reverse: 1},
			want: strings.Compare("event1", "event2"),
		},
		{
			name: "reverse order same month",
			a:    line{text: "event1\t15 Jan", key: 1, sep: "\t", reverse: -1},
			b:    line{text: "event2\t10 Jan", key: 1, sep: "\t", reverse: -1},
			want: -5,
		},
		{
			name: "reverse order different months",
			a:    line{text: "event1\t10 Feb", key: 1, sep: "\t", reverse: -1},
			b:    line{text: "event2\t15 Jan", key: 1, sep: "\t", reverse: -1},
			want: -1,
		},
		{
			name: "reverse with invalid month",
			a:    line{text: "event1\t15 Inv", key: 1, sep: "\t", reverse: -1},
			b:    line{text: "event2\t10 Jan", key: 1, sep: "\t", reverse: -1},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cmpMonth(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("cmpMonth(%q, %q) = %d, want %d", tt.a.text, tt.b.text, got, tt.want)
			}
		})
	}
}

func TestCmpHuman(t *testing.T) {
	tests := []struct {
		name string
		a, b line
		want int
	}{
		{
			name: "same suffix different values",
			a:    line{text: "file1\t10K", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "file2\t5K", key: 1, sep: "\t", reverse: 1},
			want: 5,
		},
		{
			name: "different suffixes same values",
			a:    line{text: "file1\t10M", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "file2\t10K", key: 1, sep: "\t", reverse: 1},
			want: 1,
		},
		{
			name: "different suffixes different values",
			a:    line{text: "file1\t5M", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "file2\t10K", key: 1, sep: "\t", reverse: 1},
			want: 1,
		},
		{
			name: "same value and suffix",
			a:    line{text: "file1\t10G", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "file2\t10G", key: 1, sep: "\t", reverse: 1},
			want: 0,
		},
		{
			name: "a invalid suffix",
			a:    line{text: "file1\t10X", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "file2\t5K", key: 1, sep: "\t", reverse: 1},
			want: -1,
		},
		{
			name: "b invalid suffix",
			a:    line{text: "file1\t10M", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "file2\t5N", key: 1, sep: "\t", reverse: 1},
			want: 1,
		},
		{
			name: "both invalid suffixes",
			a:    line{text: "file1\t10X", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "file2\t5Y", key: 1, sep: "\t", reverse: 1},
			want: strings.Compare("file1\t10X", "file2\t5Y"),
		},
		{
			name: "a non-numeric value",
			a:    line{text: "file1\tABCK", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "file2\t10M", key: 1, sep: "\t", reverse: 1},
			want: -1,
		},
		{
			name: "mixed case suffixes",
			a:    line{text: "file1\t10k", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "file2\t5K", key: 1, sep: "\t", reverse: 1},
			want: 5,
		},
		{
			name: "uppercase vs lowercase same value",
			a:    line{text: "file1\t10G", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "file2\t10g", key: 1, sep: "\t", reverse: 1},
			want: 0,
		},
		{
			name: "a missing column",
			a:    line{text: "file1", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "file2\t5T", key: 1, sep: "\t", reverse: 1},
			want: -1,
		},
		{
			name: "b missing column",
			a:    line{text: "file1\t2P", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "file2", key: 1, sep: "\t", reverse: 1},
			want: 1,
		},
		{
			name: "both missing columns",
			a:    line{text: "fileA", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "fileB", key: 1, sep: "\t", reverse: 1},
			want: strings.Compare("fileA", "fileB"),
		},
		{
			name: "reverse same suffix different values",
			a:    line{text: "file1\t10K", key: 1, sep: "\t", reverse: -1},
			b:    line{text: "file2\t5K", key: 1, sep: "\t", reverse: -1},
			want: -5,
		},
		{
			name: "reverse different suffixes",
			a:    line{text: "file1\t10M", key: 1, sep: "\t", reverse: -1},
			b:    line{text: "file2\t10K", key: 1, sep: "\t", reverse: -1},
			want: -1,
		},
		{
			name: "reverse invalid suffix",
			a:    line{text: "file1\t10X", key: 1, sep: "\t", reverse: -1},
			b:    line{text: "file2\t5K", key: 1, sep: "\t", reverse: -1},
			want: 1,
		},
		{
			name: "single character values",
			a:    line{text: "file1\tK", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "file2\tM", key: 1, sep: "\t", reverse: 1},
			want: -1,
		},
		{
			name: "empty values",
			a:    line{text: "file1\t", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "file2\t", key: 1, sep: "\t", reverse: 1},
			want: strings.Compare("file1", "file2"),
		},
		{
			name: "suffix without value",
			a:    line{text: "file1\tK", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "file2\t5", key: 1, sep: "\t", reverse: 1},
			want: 1,
		},
		{
			name: "value without suffix",
			a:    line{text: "file1\t10", key: 1, sep: "\t", reverse: 1},
			b:    line{text: "file2\t5K", key: 1, sep: "\t", reverse: 1},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cmpHuman(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("cmpHuman(%q, %q) = %d, want %d", tt.a.text, tt.b.text, got, tt.want)
			}
		})
	}
}
