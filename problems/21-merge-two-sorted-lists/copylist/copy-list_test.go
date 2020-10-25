package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateListNode(list []int) *ListNode {
	if len(list) == 0 {
		return nil
	}
	ln := &ListNode{list[0], nil}
	c := ln
	for i := 1; i < len(list); i++ {
		n := &ListNode{list[i], nil}
		c.Next = n
	}
	return ln
}

func PrintListNode(l *ListNode) {
	c := l
	for {
		fmt.Printf("%d ", c.Val)
		c = c.Next
	}
	fmt.Printf("\n")
}

func TestMergeTwoLists(t *testing.T) {
	t.Run("Example1", func(t *testing.T) {
		assert.True(t, true)
	})
}
