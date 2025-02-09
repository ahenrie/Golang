# Variables and Operators

## Pointers

### Creating Pointers
Pointers is a variable that stores the address of another variable.
```Go
  // variable declaration with var
	var count1 *int
	// variable with new()
	count2 := new(int)
	// Create a number to use its address
	countTemp := 4
	// Using &, create a pointer to an existing variable
	count3 := &countTemp
	// Create a pointer from some types without a temp variable
	t := &time.Time{}
```
Using the var statement, the pointer has a value of nil.

### Getting a Value From a Pointer
In order to get the value from the pointer, we need to dereference it.
Dereferencing a zero or nil pointer is a common bug in Go as the compiler can not warn us. It is best practice to always check if the pointer is nil before dereferencing it
