package tests

import (
	"reflect"
	"testing"

	hot100practice "hot100_practice"
)

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "leetcode sample 1",
			nums: []int{-4, -1, 0, 3, 10},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			name: "leetcode sample 2",
			nums: []int{-7, -3, 2, 3, 11},
			want: []int{4, 9, 9, 49, 121},
		},
		{
			name: "all negative",
			nums: []int{-5, -3, -2, -1},
			want: []int{1, 4, 9, 25},
		},
		{
			name: "all non negative",
			nums: []int{0, 1, 2, 3},
			want: []int{0, 1, 4, 9},
		},
		{
			name: "single element",
			nums: []int{-1},
			want: []int{1},
		},
		{
			name: "all zero",
			nums: []int{0, 0, 0},
			want: []int{0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hot100practice.SortedSquares(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("SortedSquares(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
