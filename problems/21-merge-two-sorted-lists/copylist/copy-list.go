package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func InitMergedList(l1, l2, l1Cursor, l2Cursor *ListNode) *ListNode {
	var l *ListNode
	if l1.Val <= l2.Val {
		l = &ListNode{l1.Val, nil}
		l1Cursor = l1.Next
	} else {
		l = &ListNode{l2.Val, nil}
		l2Cursor = l2.Next
	}
	return l
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	l1Cursor := l1
	l2Cursor := l2
	l3 := InitMergedList(l1, l2, l1Cursor, l2Cursor)
	l3Cursor := l3
	for {
		if l1Cursor.Val < l2Cursor.Val {
			l3Cursor.Next = &ListNode{l1Cursor.Val, nil}
			l3Cursor = l3Cursor.Next
			l1Cursor = l1Cursor.Next
		}
		if l1Cursor.Val > l2Cursor.Val {
			l3Cursor.Next = &ListNode{l2Cursor.Val, nil}
			l3Cursor = l3Cursor.Next
			l2Cursor = l2Cursor.Next
		}
		if l1Cursor.Val == l2Cursor.Val {
			l3Cursor.Next = &ListNode{l1Cursor.Val, nil}
			l3Cursor = l3Cursor.Next
			l1Cursor = l1Cursor.Next
			l3Cursor.Next = &ListNode{l2Cursor.Val, nil}
			l3Cursor = l3Cursor.Next
			l2Cursor = l2Cursor.Next
		}
		if l1Cursor == nil && l2Cursor == nil {
			break
		}
	}

	return l3
}
