package list

import (
	"errors"
)

var (
	ErrListFull  = errors.New("cannot insert, list is full")
	ErrListEmpty = errors.New("cannot remove, list is empty")
)

type Node struct {
	value    int64
	next     *Node
	previous *Node
}

type LinkedList struct {
	head    *Node
	tail    *Node
	count   int8
	maxSize int8
}

type ILinkedList interface {
	Destroy()
	InsertAtHead(int64) error
	InsertAtTail(int64) error
	RemoveHead() error
	RemoveTail() error
	Size() int8
	IsEmpty() bool
	IsFull() bool
	Contains(int64) bool
}

func Create() *LinkedList {
	return &LinkedList{
		head:    nil,
		tail:    nil,
		count:   0,
		maxSize: 127,
	}
}

func (ll *LinkedList) Destroy() {
	ll.head = nil
	ll.tail = nil
	ll.count = 0
}

func (ll *LinkedList) InsertAtHead(value int64) error {
	if ll.IsFull() {
		return ErrListFull
	}

	if ll.IsEmpty() {
		node := &Node{value: value}
		node.next = node
		node.previous = node
		ll.head = node
		ll.tail = node
	} else {
		node := &Node{value: value, next: ll.head, previous: ll.tail}
		ll.head.previous = node
		ll.tail.next = node
		ll.head = node
	}

	ll.count += 1
	return nil
}

func (ll *LinkedList) InsertAtTail(value int64) error {
	if ll.IsFull() {
		return ErrListFull
	}

	if ll.IsEmpty() {
		node := &Node{value: value}
		node.next = node
		node.previous = node
		ll.head = node
		ll.tail = node
	} else {
		node := &Node{value: value, next: ll.head, previous: ll.tail}
		ll.tail.next = node
		ll.head.previous = node
		ll.tail = node
	}

	ll.count += 1
	return nil
}

func (ll *LinkedList) RemoveHead() error {
	if ll.IsEmpty() {
		return ErrListEmpty
	}

	if ll.head == ll.tail {
		ll.head = nil
		ll.tail = nil
		return nil
	}

	ll.head = ll.head.next
	ll.head.previous = ll.tail
	ll.tail.next = ll.head

	return nil
}

func (ll *LinkedList) RemoveTail() error {
	if ll.IsEmpty() {
		return ErrListEmpty
	}

	if ll.head == ll.tail {
		ll.head = nil
		ll.tail = nil
		return nil
	}

	ll.tail = ll.tail.previous
	ll.tail.next = ll.head
	ll.head.previous = ll.tail

	return nil
}

func (ll *LinkedList) Head() (int64, error) {
	if ll.IsEmpty() {
		return -1, ErrListEmpty
	}

	return ll.head.value, nil
}

func (ll *LinkedList) Tail() (int64, error) {
	if ll.IsEmpty() {
		return -1, ErrListEmpty
	}

	return ll.tail.value, nil
}

func (ll *LinkedList) Size() int8 {
	return ll.maxSize
}

func (ll *LinkedList) IsEmpty() bool {
	return ll.count == 0
}

func (ll *LinkedList) IsFull() bool {
	return ll.count == ll.maxSize
}

func (ll *LinkedList) Contains(value int64) bool {
	currentNode := ll.head
	for i := 0; i < int(ll.Size()); i++ {
		if currentNode.value == value {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}
