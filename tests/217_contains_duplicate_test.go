package tests

import (
	"testing"

	hot100practice "hot100_practice"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "leetcode sample 1",
			nums: []int{1, 2, 3, 1},
			want: true,
		},
		{
			name: "leetcode sample 2",
			nums: []int{1, 2, 3, 4},
			want: false,
		},
		{
			name: "leetcode sample 3",
			nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			want: true,
		},
		{
			name: "empty input",
			nums: []int{},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hot100practice.ContainsDuplicate(tt.nums)
			if got != tt.want {
				t.Fatalf("ContainsDuplicate(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
