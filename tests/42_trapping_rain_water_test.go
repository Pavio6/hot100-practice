package tests

import (
	"testing"

	hot100practice "hot100_practice"
)

func TestTrap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "leetcode sample",
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:   6,
		},
		{
			name:   "simple valley",
			height: []int{2, 0, 2},
			want:   2,
		},
		{
			name:   "monotonic increasing",
			height: []int{1, 2, 3, 4},
			want:   0,
		},
		{
			name:   "empty",
			height: []int{},
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hot100practice.Trap(tt.height)
			if got != tt.want {
				t.Fatalf("Trap(%v) = %d, want %d", tt.height, got, tt.want)
			}
		})
	}
}

