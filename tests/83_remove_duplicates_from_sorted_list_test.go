package tests

import (
	"reflect"
	"testing"

	hot100practice "hot100_practice"
)

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{
			name: "leetcode sample 1",
			in:   []int{1, 1, 2},
			want: []int{1, 2},
		},
		{
			name: "leetcode sample 2",
			in:   []int{1, 1, 2, 3, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "single node",
			in:   []int{1},
			want: []int{1},
		},
		{
			name: "empty list",
			in:   []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildLinkedList(tt.in)
			got := hot100practice.DeleteDuplicates(head)
			if !reflect.DeepEqual(linkedListToSlice(got), tt.want) {
				t.Fatalf("DeleteDuplicates(%v) = %v, want %v", tt.in, linkedListToSlice(got), tt.want)
			}
		})
	}
}

