package tests

import (
	"reflect"
	"testing"

	hot100practice "hot100_practice"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "leetcode sample 1",
			nums: []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
		{
			name: "contains zero",
			nums: []int{-1, 1, 0, -3, 3},
			want: []int{0, 0, 9, 0, 0},
		},
		{
			name: "two elements",
			nums: []int{5, 10},
			want: []int{10, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hot100practice.ProductExceptSelf(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("ProductExceptSelf(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

