package tests

import (
	"reflect"
	"sort"
	"strings"
	"testing"

	hot100practice "hot100_practice"
)

func normalizeGroups(groups [][]string) [][]string {
	normalized := make([][]string, len(groups))
	for i, g := range groups {
		cp := append([]string(nil), g...)
		sort.Strings(cp)
		normalized[i] = cp
	}

	sort.Slice(normalized, func(i, j int) bool {
		return strings.Join(normalized[i], ",") < strings.Join(normalized[j], ",")
	})
	return normalized
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		want [][]string
	}{
		{
			name: "leetcode sample",
			in:   []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}},
		},
		{
			name: "single empty string",
			in:   []string{""},
			want: [][]string{{""}},
		},
		{
			name: "single char",
			in:   []string{"a"},
			want: [][]string{{"a"}},
		},
		{
			name: "empty input",
			in:   []string{},
			want: [][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hot100practice.GroupAnagrams(tt.in)
			if !reflect.DeepEqual(normalizeGroups(got), normalizeGroups(tt.want)) {
				t.Fatalf("GroupAnagrams(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}
