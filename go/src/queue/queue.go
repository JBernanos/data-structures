package queue

import (
	"errors"
)

var (
	ErrQueueFull  = errors.New("cannot enqueue, queue is full")
	ErrQueueEmpty = errors.New("cannot dequeue/peek, queue is empty")
)

type Node struct {
	value int64
	next  *Node
}

type Queue struct {
	front   *Node
	count   int8
	maxSize int8
}

type IQueue interface {
	Destroy()
	Enqueue(value int64) error
	Dequeue() error
	Peek() (int64, error)
	Size() int8
	IsEmpty() bool
	IsFull() bool
	Contains(int64) bool
}

func Create() *Queue {
	return &Queue{
		front:   nil,
		count:   0,
		maxSize: 127,
	}
}

func (q *Queue) Destroy() {
	q.front = nil
	q.count = 0
}

func (q *Queue) Enqueue(value int64) error {
	if q.IsFull() {
		return ErrQueueFull
	}

	node := &Node{value: value, next: nil}

	if q.IsEmpty() {
		q.front = node
		q.count += 1
		return nil
	}

	currentNode := q.front
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = node
	q.count += 1

	return nil
}

func (q *Queue) Dequeue() error {
	if q.IsEmpty() {
		return ErrQueueEmpty
	}

	q.front = q.front.next
	q.count -= 1

	return nil
}

func (q *Queue) Peek() (int64, error) {
	if q.IsEmpty() {
		return -1, ErrQueueEmpty
	}

	return q.front.value, nil
}

func (q *Queue) Size() int8 {
	return q.maxSize
}

func (q *Queue) IsEmpty() bool {
	return q.count == 0
}

func (q *Queue) IsFull() bool {
	return q.count == q.maxSize
}

func (q *Queue) Contains(value int64) bool {
	currentNode := q.front
	for currentNode != nil {
		if currentNode.value == value {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}
