package main

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	l1 := CreateListNode([]int{2, 4, 3})
	l2 := CreateListNode([]int{5, 6, 4})
	PrintListNode(l1)
	PrintListNode(l2)
	addTwoNumbers(l1, l2)
}

func BenchmarkAddTwoNumbers(b *testing.B) {
	b.ResetTimer()
	var l1 *ListNode
	var l2 *ListNode
	for i := 0; i < b.N; i++ {
		addTwoNumbers(l1, l2)
	}
}
