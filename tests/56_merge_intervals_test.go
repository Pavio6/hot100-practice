package tests

import (
	"reflect"
	"testing"

	hot100practice "hot100_practice"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      [][]int
	}{
		{
			name:      "leetcode sample 1",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			want:      [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "leetcode sample 2",
			intervals: [][]int{{1, 4}, {4, 5}},
			want:      [][]int{{1, 5}},
		},
		{
			name:      "contains interval",
			intervals: [][]int{{1, 4}, {2, 3}},
			want:      [][]int{{1, 4}},
		},
		{
			name:      "unsorted input",
			intervals: [][]int{{8, 10}, {1, 3}, {2, 6}},
			want:      [][]int{{1, 6}, {8, 10}},
		},
		{
			name:      "empty input",
			intervals: [][]int{},
			want:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hot100practice.Merge(tt.intervals)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("Merge(%v) = %v, want %v", tt.intervals, got, tt.want)
			}
		})
	}
}
