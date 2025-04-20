package main

import (
	"data-structures/src/stack"
	"fmt"
)

func main() {
	fmt.Println("Data Structures implementation in GO.")

	// Stack
	var s stack.IStack = stack.Create()
	s.Push(1)
}
