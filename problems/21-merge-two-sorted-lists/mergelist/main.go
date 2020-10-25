package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func InitMergedList(l1, l2 *ListNode) (*ListNode, *ListNode, *ListNode) {
	if l1.Val <= l2.Val {
		return l1, l1.Next, l2
	} else {
		return l2, l1, l2.Next
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	l3, c1, c2 := InitMergedList(l1, l2)
	c3 := l3
	for {
		if c1 == nil {
			c3.Next = c2
			return l3
		}
		if c2 == nil {
			c3.Next = c1
			return l3
		}
		if c1.Val <= c2.Val {
			c3.Next = c1
			c1 = c1.Next
			c3 = c3.Next
		} else {
			// c1.Val > c2.Val
			c3.Next = c2
			c2 = c2.Next
			c3 = c3.Next
		}
		if c1 == nil && c2 == nil {
			break
		}
	}

	return l3
}
