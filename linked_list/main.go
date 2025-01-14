package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
}

func main() {
	var myList linkedList
	myList.addNode(1)
	myList.addNode(2)
	myList.addNode(3)
	myList.addNode(4)
	myList.addNode(5)

	myList.printList()

}

func (list *linkedList) addNode(data int) {
	newNode := &node{data, nil}

	if list.head == nil {
		list.head = newNode
	} else {
		// Start at the beginning
		currentNode := list.head

		// Loop through the list till we get to the end of the list
		for currentNode.next != nil {
			// Go to the next node
			currentNode = currentNode.next
		}
		// Now we are at the end
		currentNode.next = newNode
	}
}

func (list *linkedList) printList() {
	currentNode := list.head
	for currentNode.next != nil {
		fmt.Printf("Current Node Address: %p \t Node Data: %d \t Next Node: %p\n", currentNode, currentNode.data, currentNode.next)
		currentNode = currentNode.next
	}
}
