package main

import (
	"fmt"
	"math/big"
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
	l := CreateListNode([]int{1, 0, 1})
	//PrintListNode(l)
	return l
}

func addBigInt() {
	new(big.Int).Add()
}
