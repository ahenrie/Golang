/*
Testing for pointers
*/

package main

import (
	"fmt"

	pointers "github.com/ahenrie/VarAndOp/Pointers"
)

func main() {
	fmt.Println("hello world")
	pointer1 := pointers.Allocate(5)
	pointer2 := pointers.Allocate(6)
	pointers.Deref(pointer1)
	pointers.Deref(pointer2)

	myList, err := pointers.BuildList(100)

	if err == nil {
		myList.PrintList()
	} else {
		fmt.Println(err)
	}

}
