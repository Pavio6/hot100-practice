package tests

import (
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"

	hot100practice "hot100_practice"
)

func normalizeTriplets(in [][]int) [][]int {
	out := make([][]int, len(in))
	for i, triplet := range in {
		cp := append([]int(nil), triplet...)
		sort.Ints(cp)
		out[i] = cp
	}

	sort.Slice(out, func(i, j int) bool {
		return tripletKey(out[i]) < tripletKey(out[j])
	})
	return out
}

func tripletKey(t []int) string {
	parts := make([]string, len(t))
	for i, v := range t {
		parts[i] = strconv.Itoa(v)
	}
	return strings.Join(parts, ",")
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "leetcode sample",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "all zeros",
			nums: []int{0, 0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "no solution",
			nums: []int{1, 2, -2, -1},
			want: [][]int{},
		},
		{
			name: "empty input",
			nums: []int{},
			want: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hot100practice.ThreeSum(tt.nums)
			if !reflect.DeepEqual(normalizeTriplets(got), normalizeTriplets(tt.want)) {
				t.Fatalf("ThreeSum(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
