package main

import "testing"

func TestReplaceBit(t *testing.T) {
	tests := []struct {
		name     string
		num      int64
		bit      int64
		shift    uint8
		expected int64
	}{
		{"Set first bit to 1", 0b0, 1, 1, 0b1},
		{"Clear second bit", 0b11, 0, 2, 0b01},
		{"Set third bit in zero", 0, 1, 3, 0b100},
		{"No change when same value (0)", 0b1010, 0, 3, 0b1010},
		{"No change when same value (1)", 0b1010, 1, 2, 0b1010},
		{"Set last (64th) bit", 0, 1, 64, -9223372036854775808},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReplaceBit(tt.num, tt.bit, tt.shift)
			if got != tt.expected {
				t.Errorf("ReplaceBit(%d, %d, %d) = %d (%b), want %d (%b)",
					tt.num, tt.bit, tt.shift, got, got, tt.expected, tt.expected)
			}
		})
	}
}
