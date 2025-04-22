package deque

import (
	"errors"
)

var (
	ErrDequeFull  = errors.New("cannot insert, deque is full")
	ErrDequeEmpty = errors.New("cannot remove/peek, deque is empty")
)

type Node struct {
	value    int64
	next     *Node
	previous *Node
}

type Deque struct {
	front   *Node
	back    *Node
	count   int8
	maxSize int8
}

type IDeque interface {
	Destroy()
	InsertFront(int64) error
	InsertBack(int64) error
	RemoveFront() error
	RemoveBack() error
	Size() int8
	IsEmpty() bool
	IsFull() bool
	Contains(int64) bool
}

func Create() *Deque {
	return &Deque{
		front:   nil,
		back:    nil,
		count:   0,
		maxSize: 127,
	}
}

func (d *Deque) Destroy() {
	d.front = nil
	d.back = nil
	d.count = 0
}

func (d *Deque) InsertFront(value int64) error {
	if d.IsFull() {
		return ErrDequeFull
	}

	node := &Node{value: value, next: d.front, previous: nil}

	if d.IsEmpty() {
		d.back = node
	} else {
		d.front.previous = node
	}
	d.front = node
	d.count += 1

	return nil
}

func (d *Deque) InsertBack(value int64) error {
	if d.IsFull() {
		return ErrDequeFull
	}

	node := &Node{value: value, next: nil, previous: d.back}
	if d.IsEmpty() {
		d.front = node
	} else {
		d.back.next = node
	}
	d.back = node
	d.count += 1

	return nil
}

func (d *Deque) RemoveFront() error {
	if d.IsEmpty() {
		return ErrDequeEmpty
	}

	d.front = d.front.next
	if d.front != nil {
		d.front.previous = nil
	} else {
		d.back = nil
	}
	d.count -= 1

	return nil
}

func (d *Deque) RemoveBack() error {
	if d.IsEmpty() {
		return ErrDequeEmpty
	}

	d.back = d.back.previous
	if d.back != nil {
		d.back.next = nil
	} else {
		d.front = nil
	}
	d.count -= 1

	return nil
}

func (d *Deque) Size() int8 {
	return d.maxSize
}

func (d *Deque) IsEmpty() bool {
	return d.count == 0
}

func (d *Deque) IsFull() bool {
	return d.count == d.maxSize
}

func (d *Deque) Contains(value int64) bool {
	currentNode := d.front
	for currentNode != nil {
		if currentNode.value == value {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}
