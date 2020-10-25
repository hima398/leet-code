package main

import (
	"fmt"
	"strings"
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
		c = c.Next
	}
	return ln
}

func FormatListNode(l *ListNode) string {
	var ss []string
	c := l
	for {
		if c == nil {
			break
		}
		ss = append(ss, fmt.Sprintf("%d", c.Val))
		c = c.Next
	}
	return "[" + strings.Join(ss, ",") + "]"
}

func TestMergeTwoLists(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		l1 := CreateListNode([]int{1, 2, 4})
		l2 := CreateListNode([]int{1, 3, 4})
		l3 := mergeTwoLists(l1, l2)
		assert.Equal(t, "[1,1,2,3,4,4]", FormatListNode(l3))
	})
	t.Run("Example 2", func(t *testing.T) {
		l1 := CreateListNode([]int{})
		l2 := CreateListNode([]int{})
		l3 := mergeTwoLists(l1, l2)
		assert.Equal(t, "[]", FormatListNode(l3))
	})
	t.Run("Example 3", func(t *testing.T) {
		l1 := CreateListNode([]int{})
		l2 := CreateListNode([]int{0})
		l3 := mergeTwoLists(l1, l2)
		assert.Equal(t, "[0]", FormatListNode(l3))
	})
	t.Run("Example 4", func(t *testing.T) {
		l1 := CreateListNode([]int{0})
		l2 := CreateListNode([]int{})
		l3 := mergeTwoLists(l1, l2)
		assert.Equal(t, "[0]", FormatListNode(l3))
	})
}
