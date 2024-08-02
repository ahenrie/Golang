package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Print to the console and ask what would like to be screamed.
	fmt.Println("What would you like me to scream?")

	// Allocate a variable named in and use a Reader from the bufio library that reads
	// standard input from the user
	in := bufio.NewReader(os.Stdin)

	// Allocate a variable called s and ignore the error return value "_"
	// Store what the Reader objet reads until the newline character
	s, _ := in.ReadString('\n')

	// Trim whitespace on the input
	s = strings.TrimSpace(s)

	// Force it to uppercase
	s = strings.ToUpper(s)

	// Print the s variable to the console.
	fmt.Println(s)
}
