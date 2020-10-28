package main

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		l1 := CreateListNode([]int{2, 4, 3})
		l2 := CreateListNode([]int{5, 6, 4})
		PrintListNode(l1)
		PrintListNode(l2)
		l := addTwoNumbers(l1, l2)
		PrintListNode(l)
	})
	t.Run("Example 2", func(t *testing.T) {
		l1 := CreateListNode([]int{0})
		l2 := CreateListNode([]int{0})
		PrintListNode(l1)
		PrintListNode(l2)
		l := addTwoNumbers(l1, l2)
		PrintListNode(l)
	})
	t.Run("Example 3", func(t *testing.T) {
		l1 := CreateListNode([]int{9, 9, 9, 9, 9, 9, 9})
		l2 := CreateListNode([]int{9, 9, 9, 9})
		PrintListNode(l1)
		PrintListNode(l2)
		l := addTwoNumbers(l1, l2)
		PrintListNode(l)
	})
}

func BenchmarkAddTwoNumbers(b *testing.B) {
	b.ResetTimer()
	var l1 *ListNode
	var l2 *ListNode
	for i := 0; i < b.N; i++ {
		addTwoNumbers(l1, l2)
	}
}
