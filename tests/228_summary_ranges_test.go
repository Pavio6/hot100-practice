package tests

import (
	"reflect"
	"testing"

	hot100practice "hot100_practice"
)

func TestSummaryRanges(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []string
	}{
		{
			name: "leetcode sample 1",
			nums: []int{0, 1, 2, 4, 5, 7},
			want: []string{"0->2", "4->5", "7"},
		},
		{
			name: "leetcode sample 2",
			nums: []int{0, 2, 3, 4, 6, 8, 9},
			want: []string{"0", "2->4", "6", "8->9"},
		},
		{
			name: "single element",
			nums: []int{5},
			want: []string{"5"},
		},
		{
			name: "empty input",
			nums: []int{},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hot100practice.SummaryRanges(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("SummaryRanges(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

