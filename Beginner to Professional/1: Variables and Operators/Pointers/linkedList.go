package pointers

import (
	"errors"
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
}

// Constructor
func newList() *linkedList {
	return &linkedList{}
}

func (linkList *linkedList) insert(val int) {
	// create node to hold the new number
	newNode := &node{val, nil}

	// if empty
	if linkList.head == nil {
		linkList.head = newNode
	} else {
		//iterate
		currentNode := linkList.head

		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
}

func BuildList(num int) (*linkedList, error) {
	// build a linked list from 0 to num

	//check num > 0
	if num < 0 {
		blank := newList()
		return blank, errors.New("Lets use a positive number okay!")
	}

	// while i < num, create a node and add it to the linkedList
	myList := newList()
	for i := 0; i <= num; i++ {
		//create node then insert
		myList.insert(i)
	}
	return myList, nil
}

func (linkList *linkedList) PrintList() {
	// Check for empty list
	if linkList.head == nil {
		fmt.Println("Your list is empty! Try adding some nodes first!")
	} else {
		// Iterate through the list starting at the head
		currentNode := linkList.head
		counter := 0
		for currentNode != nil {
			fmt.Printf("Node #: %d at Address: %p --> Contains: %d\n", counter, currentNode, currentNode.data)
			counter++
			currentNode = currentNode.next
		}
	}
}
