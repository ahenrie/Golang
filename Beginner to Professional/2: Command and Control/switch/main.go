package main

import (
	"fmt"
	"time"
)

func main() {
	dayBorn := time.Monday

	switch dayBorn {
	case time.Monday:
		fmt.Println("Today is monday")
	case time.Tuesday:
		fmt.Println("Well today is not monday, its tuesday")
	default:
		fmt.Println("Error in day")
	}
}
