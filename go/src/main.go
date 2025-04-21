package main

import (
	"data-structures/src/queue"
	"data-structures/src/stack"
	"fmt"
)

func main() {
	fmt.Println("Data Structures implementation in GO.")

	// Stack
	var s stack.IStack = stack.Create()
	s.Push(1)

	// Queue
	var q queue.IQueue = queue.Create()
	q.Enqueue(1)
}
