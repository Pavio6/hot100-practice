package tests

import (
	"reflect"
	"testing"

	hot100practice "hot100_practice"
)

func TestDeleteDuplicates02(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{
			name: "leetcode sample",
			in:   []int{1, 2, 3, 3, 4, 4, 5},
			want: []int{1, 2, 5},
		},
		{
			name: "duplicates at head",
			in:   []int{1, 1, 1, 2, 3},
			want: []int{2, 3},
		},
		{
			name: "all duplicates",
			in:   []int{1, 1},
			want: []int{},
		},
		{
			name: "duplicates at tail",
			in:   []int{1, 2, 2},
			want: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildLinkedList(tt.in)
			got := hot100practice.DeleteDuplicates02(head)
			if !reflect.DeepEqual(linkedListToSlice(got), tt.want) {
				t.Fatalf("DeleteDuplicates02(%v) = %v, want %v", tt.in, linkedListToSlice(got), tt.want)
			}
		})
	}
}

