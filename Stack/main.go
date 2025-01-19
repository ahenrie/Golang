package main

import "fmt"

type stack struct {
	data []int
}

func (inStack *stack) push(inInt int) {
	inStack.data = append(inStack.data, inInt)
}

func (inStack *stack) pop() int {
	// Ensure there is something to pop
	if len(inStack.data) == 0 {
		fmt.Println("pop from empty stack") // Handle underflow case
		return -1
	}

	lastIndex := len(inStack.data) - 1
	lastInt := inStack.data[lastIndex]
	inStack.data = inStack.data[:lastIndex] // Update stack to exclude last element
	return lastInt
}

func main() {
	newStack := &stack{}
	newStack.push(1)
	newStack.push(2)
	newStack.push(3)
	fmt.Println(newStack)
	fmt.Println(newStack.pop())
	fmt.Println(newStack.pop())
	fmt.Println(newStack.pop())
	fmt.Println(newStack.pop())

}
