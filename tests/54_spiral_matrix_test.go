package tests

import (
	"reflect"
	"testing"

	hot100practice "hot100_practice"
)

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   []int
	}{
		{
			name:   "leetcode sample 3x3",
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:   []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name:   "leetcode sample 3x4",
			matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			want:   []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			name:   "single row",
			matrix: [][]int{{1, 2, 3}},
			want:   []int{1, 2, 3},
		},
		{
			name:   "single column",
			matrix: [][]int{{1}, {2}, {3}},
			want:   []int{1, 2, 3},
		},
		{
			name:   "empty matrix",
			matrix: [][]int{},
			want:   []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hot100practice.SpiralOrder(tt.matrix)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("SpiralOrder(%v) = %v, want %v", tt.matrix, got, tt.want)
			}
		})
	}
}
