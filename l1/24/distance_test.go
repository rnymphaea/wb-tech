package main

import (
	"math"
	"testing"
)

func TestNewPoint(t *testing.T) {
	tests := []struct {
		name string
		x    float64
		y    float64
	}{
		{"Zero point", 0, 0},
		{"Positive coordinates", 5.5, 10.2},
		{"Negative coordinates", -3.2, -7.8},
		{"Mixed coordinates", -4.5, 9.1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPoint(tt.x, tt.y)
			if p.x != tt.x || p.y != tt.y {
				t.Errorf("NewPoint() = (%v, %v), want (%v, %v)", p.x, p.y, tt.x, tt.y)
			}
		})
	}
}

func TestDistance(t *testing.T) {
	tests := []struct {
		name  string
		p1    *Point
		p2    *Point
		want  float64
		delta float64 // допустимая погрешность
	}{
		{
			name:  "Same point",
			p1:    NewPoint(0, 0),
			p2:    NewPoint(0, 0),
			want:  0,
			delta: 0.0001,
		},
		{
			name:  "Horizontal distance",
			p1:    NewPoint(0, 0),
			p2:    NewPoint(5, 0),
			want:  5,
			delta: 0.0001,
		},
		{
			name:  "Vertical distance",
			p1:    NewPoint(0, 0),
			p2:    NewPoint(0, 3),
			want:  3,
			delta: 0.0001,
		},
		{
			name:  "Diagonal distance",
			p1:    NewPoint(1, 1),
			p2:    NewPoint(4, 5),
			want:  5,
			delta: 0.0001,
		},
		{
			name:  "Negative coordinates",
			p1:    NewPoint(-2, -3),
			p2:    NewPoint(-5, -7),
			want:  5,
			delta: 0.0001,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p1.Distance(tt.p2)
			if math.Abs(got-tt.want) > tt.delta {
				t.Errorf("Distance() = %v, want %v (delta %v)", got, tt.want, tt.delta)
			}
		})
	}
}
