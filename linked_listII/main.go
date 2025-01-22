package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
}

func newLinkedList() *linkedList {
	return &linkedList{}
}

func main() {
	myList := newLinkedList()
	myList.insert(13)
	myList.insert(232)
	myList.insert(3333)
	myList.insert(422)
	myList.insert(52222)
	myList.printList()

	fmt.Println("")
	myList.reverse()
	myList.printList()
}

func (linkList *linkedList) insert(val int) {
	// New node to add to the end of the list, that we pass the passed value to when instantiating it
	newNode := &node{val, nil}

	// Checks if the list is empty
	if linkList.head == nil {
		// The head points to the new node
		linkList.head = newNode
	} else {
		// Iterate through until the node.next points to nil (we are at the end)
		currentNode := linkList.head

		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		// Now we are at the end
		currentNode.next = newNode
	}
}

func (linkList *linkedList) printList() {
	// Check for empty list
	if linkList.head == nil {
		fmt.Println("Your list is empty! Try adding some nodes first!")
	} else {
		// Iterate through the list starting at the head
		currentNode := linkList.head
		counter := 0
		for currentNode != nil {
			fmt.Printf("Node #:%d --> Contains: %d\n", counter, currentNode.data)
			counter++
			currentNode = currentNode.next
		}
	}
}

func (linkList *linkedList) reverse() {
	// Start at the beginning
	var prev *node = nil
	currentNode := linkList.head

	//Check for empty linkList
	if currentNode.next == nil {
		fmt.Println("Your list is empty!")
	} else {
		for currentNode != nil {
			// The next node = next
			nextNode := currentNode.next
			currentNode.next = prev
			prev = currentNode
			currentNode = nextNode
		}
		linkList.head = prev
	}
}
