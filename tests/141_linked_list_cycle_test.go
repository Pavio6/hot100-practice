package tests

import (
	"testing"

	hot100practice "hot100_practice"
)

func TestHasCycle(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		if hot100practice.HasCycle(nil) {
			t.Fatalf("HasCycle(nil) = true, want false")
		}
	})

	t.Run("single node no cycle", func(t *testing.T) {
		head := &hot100practice.ListNode{Val: 1}
		if hot100practice.HasCycle(head) {
			t.Fatalf("HasCycle(single no cycle) = true, want false")
		}
	})

	t.Run("single node self cycle", func(t *testing.T) {
		head := &hot100practice.ListNode{Val: 1}
		head.Next = head
		if !hot100practice.HasCycle(head) {
			t.Fatalf("HasCycle(single self cycle) = false, want true")
		}
	})

	t.Run("multi node with cycle", func(t *testing.T) {
		head := buildLinkedList([]int{3, 2, 0, -4})
		tail := head
		for tail.Next != nil {
			tail = tail.Next
		}
		tail.Next = head.Next
		if !hot100practice.HasCycle(head) {
			t.Fatalf("HasCycle(multi cycle) = false, want true")
		}
	})

	t.Run("multi node no cycle", func(t *testing.T) {
		head := buildLinkedList([]int{1, 2, 3})
		if hot100practice.HasCycle(head) {
			t.Fatalf("HasCycle(multi no cycle) = true, want false")
		}
	})
}

