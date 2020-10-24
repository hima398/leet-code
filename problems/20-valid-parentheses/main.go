package main

var pMap = map[uint8]uint8{
	40:  41,  // ()
	123: 125, // {}
	91:  93,  // []
}

type Stack struct {
	top    *Node
	length int
}

type Node struct {
	value uint8
	prev  *Node
}

func NewStack() *Stack {
	return &Stack{nil, 0}
}

func (this *Stack) Pop() *Node {
	if this.length <= 0 {
		return nil
	}

	n := this.top
	this.top = n.prev
	this.length--
	return n
}

func (this *Stack) Push(v uint8) {
	n := &Node{v, this.top}
	this.top = n
	this.length++
}

func isValid(s string) bool {
	stack := NewStack()
	for i := 0; i < len(s); i++ {
		// Close Parenthese
		if stack.length <= 0 && (s[i] == 41 || s[i] == 125 || s[i] == 93) {
			return false
		}
		if stack.length > 0 {
			open := stack.top.value
			if pMap[open] == s[i] {
				stack.Pop()
				continue
			}
			if open == 40 && s[i] == 125 || open == 40 && s[i] == 93 || open == 123 && s[i] == 41 || open == 123 && s[i] == 93 || (open == 91 && s[i] == 41) || (open == 91 && s[i] == 125) {
				return false
			}
		}
		// Open Parenthse
		if s[i] == 40 || s[i] == 123 || s[i] == 91 {
			stack.Push(s[i])
			continue
		}
	}
	return stack.length == 0
}
