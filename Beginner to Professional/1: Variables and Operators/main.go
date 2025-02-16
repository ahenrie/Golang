/*
Testing for pointers
*/

package main

import (
	"fmt"

	constants "github.com/ahenrie/VarAndOp/Constants"
	pointers "github.com/ahenrie/VarAndOp/Pointers"
)

func main() {
	fmt.Println("hello world")
	pointer1 := pointers.Allocate(5)
	pointer2 := pointers.Allocate(6)
	pointers.Deref(pointer1)
	pointers.Deref(pointer2)

	myList, err := pointers.BuildList(10)

	if err == nil {
		myList.PrintList()
	} else {
		fmt.Println(err)
	}

	// Adding to a pointer with a function
	pointers.Add(100, pointer1)

	a, b := 5, 10
	pointers.Swap(&a, &b)
	fmt.Println(a == 10 && b == 5)

	myCache := constants.Cache{}

	myCache.SetBook("1234-5678", "Get Ready To Go")
	myCache.SetCD("1234-5678", "Get Ready To Go Audio Book")

	fmt.Println("Book : ", myCache.GetBook("1234-5678"))
	fmt.Println("CD : ", myCache.GetCD("1234-5678"))

}
