package tests

import (
	"reflect"
	"testing"

	hot100practice "hot100_practice"
)

func TestMiddleNode(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{
			name: "odd length",
			in:   []int{1, 2, 3, 4, 5},
			want: []int{3, 4, 5},
		},
		{
			name: "even length return second middle",
			in:   []int{1, 2, 3, 4, 5, 6},
			want: []int{4, 5, 6},
		},
		{
			name: "single node",
			in:   []int{1},
			want: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildLinkedList(tt.in)
			got := hot100practice.MiddleNode(head)
			if !reflect.DeepEqual(linkedListToSlice(got), tt.want) {
				t.Fatalf("MiddleNode(%v) = %v, want %v", tt.in, linkedListToSlice(got), tt.want)
			}
		})
	}
}

