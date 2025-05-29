package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := ""

	fmt.Print("Enter a number: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error reading input: ", err)
		return
	}
	number, err := strconv.Atoi(input)
	fmt.Println("You entered: ", number)

	for i := 1; i <= number; i++ {
		out_string := ""

		if i%3 == 0 {
			out_string += "Fizz"
		}

		if i%5 == 0 {
			out_string += "Buzz"
		}

		if out_string == "" {
			out_string = strconv.Itoa(i)
		}

		fmt.Println(out_string)
	}
}
