package tests

import (
	"reflect"
	"testing"

	hot100practice "hot100_practice"
)

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		n    int
		want []int
	}{
		{
			name: "leetcode sample",
			in:   []int{1, 2, 3, 4, 5},
			n:    2,
			want: []int{1, 2, 3, 5},
		},
		{
			name: "single node",
			in:   []int{1},
			n:    1,
			want: []int{},
		},
		{
			name: "remove tail",
			in:   []int{1, 2},
			n:    1,
			want: []int{1},
		},
		{
			name: "remove head",
			in:   []int{1, 2},
			n:    2,
			want: []int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildLinkedList(tt.in)
			got := hot100practice.RemoveNthFromEnd(head, tt.n)
			if !reflect.DeepEqual(linkedListToSlice(got), tt.want) {
				t.Fatalf("RemoveNthFromEnd(%v, %d) = %v, want %v", tt.in, tt.n, linkedListToSlice(got), tt.want)
			}
		})
	}
}

