package stack

import (
	"errors"
)

var (
	ErrStackFull  = errors.New("cannot push, stack is full")
	ErrStackEmpty = errors.New("cannot pop/peek, stack is empty")
)

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
	Push(value int64) error
	Pop() error
	Peek() (int64, error)
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

func (s *Stack) Push(value int64) error {
	if s.IsFull() {
		return ErrStackFull
	}

	node := &Node{value: value, next: s.top}
	s.top = node
	s.count += 1
	return nil
}

func (s *Stack) Pop() error {
	if s.IsEmpty() {
		return ErrStackEmpty
	}

	s.top = s.top.next
	s.count -= 1
	return nil
}

func (s *Stack) Peek() (int64, error) {
	if s.IsEmpty() {
		return -1, ErrStackEmpty
	}

	return s.top.value, nil
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
