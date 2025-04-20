package stack

import "fmt"

type Node struct {
	value int64
	next  *Node
}

type Stack struct {
	top     *Node
	count   int8
	maxSize int8
}

type IStack interface {
	Destroy()
	Push(value int64)
	Pop()
	Peek() int64
	Size() int8
	IsEmpty() bool
	IsFull() bool
	Contains(int64) bool
}

func Create() *Stack {
	return &Stack{
		top:     nil,
		count:   0,
		maxSize: 127,
	}
}

func (s *Stack) Destroy() {
	s.top = nil
	s.count = 0
}

func (s *Stack) Push(value int64) {
	if s.IsFull() {
		fmt.Printf("(ERROR) - Stack max size (%d) reached. \n", s.maxSize)
		return
	}

	node := Node{value: value, next: s.top}
	s.top = &node
	s.count += 1
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		fmt.Println("(ERROR) - Cannot pop, stack is empty.")
		return
	}

	s.top = s.top.next
	s.count -= 1
}

func (s *Stack) Peek() int64 {
	return s.top.value
}

func (s *Stack) Size() int8 {
	return s.maxSize
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) IsFull() bool {
	return s.count == s.maxSize
}

func (s *Stack) Contains(value int64) bool {
	currentNode := s.top
	for currentNode != nil {
		if currentNode.value == value {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}
