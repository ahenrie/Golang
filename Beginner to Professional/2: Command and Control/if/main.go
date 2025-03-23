package main

/*
An if is the most basic form of logic
*/

import "fmt"


func main() {
  input := -10
  
  if input < 0 {
    fmt.Println("Input Can Not Be Negative.")
  } else if input%2 == 0 {
    fmt.Println(input, " is even.")
  } else {
    fmt.Println(input, " is odd.")
  }
}
