package tests

import (
	"reflect"
	"testing"

	hot100practice "hot100_practice"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "basic",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "duplicate values",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "no answer",
			nums:   []int{1, 2, 3},
			target: 10,
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hot100practice.TwoSum03(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("TwoSum03(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
