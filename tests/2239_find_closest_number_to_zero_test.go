package tests

import (
	"testing"

	hot100practice "hot100_practice"
)

func TestFindClosestNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "leetcode sample",
			nums: []int{-4, -2, 1, 4, 8},
			want: 1,
		},
		{
			name: "tie choose positive",
			nums: []int{-2, 2},
			want: 2,
		},
		{
			name: "all negative",
			nums: []int{-7, -5, -3, -1},
			want: -1,
		},
		{
			name: "contains zero",
			nums: []int{-10, 0, 5},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hot100practice.FindClosestNumber(tt.nums)
			if got != tt.want {
				t.Fatalf("FindClosestNumber(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
