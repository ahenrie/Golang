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
		default:
			fmt.Println("Error in day")
	} 
}
