package tests

import (
	"testing"

	hot100practice "hot100_practice"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "leetcode sample 1",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			name:   "leetcode sample 2",
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		{
			name:   "single day",
			prices: []int{5},
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hot100practice.MaxProfit(tt.prices)
			if got != tt.want {
				t.Fatalf("MaxProfit(%v) = %d, want %d", tt.prices, got, tt.want)
			}
		})
	}
}

