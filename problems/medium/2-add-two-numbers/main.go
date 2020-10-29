package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintListNode(ln *ListNode) {
	for c := ln; c != nil; c = c.Next {
		fmt.Println(c)
	}
}

func CreateListNode(l []int) *ListNode {
	n := len(l)
	// 1 <= len(l) <= 100
	ln := &ListNode{l[0], nil}
	c := ln
	for i := 1; i < n; i++ {
		node := &ListNode{l[i], nil}
		c.Next = node
		c = node
	}
	return ln
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ln := &ListNode{}
	c := ln
	c1 := l1
	c2 := l2
	carry := 0
	for c1 != nil || c2 != nil || carry != 0 {
		v := carry
		if c1 != nil {
			v += c1.Val
			c1 = c1.Next
		}
		if c2 != nil {
			v += c2.Val
			c2 = c2.Next
		}
		if v >= 10 {
			v %= 10
			carry = 1
		} else {
			carry = 0
		}
		c.Next = &ListNode{v, nil}
		c = c.Next
	}
	return ln.Next
}
