package tests

import (
	"testing"

	hot100practice "hot100_practice"
)

func TestCalPoints(t *testing.T) {
	tests := []struct {
		name       string
		operations []string
		want       int
	}{
		{
			name:       "leetcode sample 1",
			operations: []string{"5", "2", "C", "D", "+"},
			want:       30,
		},
		{
			name:       "leetcode sample 2",
			operations: []string{"5", "-2", "4", "C", "D", "9", "+", "+"},
			want:       27,
		},
		{
			name:       "all numeric",
			operations: []string{"1", "2", "3"},
			want:       6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hot100practice.CalPoints(tt.operations)
			if got != tt.want {
				t.Fatalf("CalPoints(%v) = %d, want %d", tt.operations, got, tt.want)
			}
		})
	}
}
