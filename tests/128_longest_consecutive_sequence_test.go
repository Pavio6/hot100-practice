package tests

import (
	"testing"

	hot100practice "hot100_practice"
)

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "leetcode sample",
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
		{
			name: "unordered long chain with duplicates",
			nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want: 9,
		},
		{
			name: "empty input",
			nums: []int{},
			want: 0,
		},
		{
			name: "single element",
			nums: []int{5},
			want: 1,
		},
		{
			name: "duplicates mixed in",
			nums: []int{1, 2, 2, 3},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hot100practice.LongestConsecutive(tt.nums)
			if got != tt.want {
				t.Fatalf("LongestConsecutive(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
