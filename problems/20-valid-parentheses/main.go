package main

import "fmt"

var pMap = map[byte]byte{
	40:  41,  // ()
	123: 125, // {}
	91:  93,  // []
}

type Stack struct {
	top    *Node
	length int
}

type Node struct {
	value byte
	prev  *Node
}

func NewStack() *Stack {
	return &Stack{nil, 0}
}

func (this *Stack) Pop() *Node {
	fmt.Println("Call Pop")
	if this.length <= 0 {
		return nil
	}

	n := this.top
	this.top = n.prev
	this.length--
	return n
}

func (this *Stack) Push(v byte) {
	fmt.Println("Call Push")
	n := &Node{v, this.top}
	this.top = n
	this.length++
}

func isValid(s string) bool {
	stack := NewStack()
	for i := 0; i < len(s); i++ {
		// Open Parenthse
		if s[i] == 40 || s[i] == 123 || s[i] == 91 {
			stack.Push(s[i])
		}
		// Close Parenthese
		if stack.length > 0 && stack.top.value == pMap[s[i]] {
			stack.Pop()
		}
		fmt.Println(stack)
	}
	return stack.length == 0
}
