package main

import "fmt"
//we will get user input for a number then build a slice from 0 to the input 

//var int inputNumber
//fmt.Println("Enter a number for array size:")
//fmt.Scanln(&inputNumber)

func get_number() int {
  //keep asking for int until int > 0
  var num_int int
  for {
    fmt.Println("Enter a number greater than 0:")
    fmt.Scanln(&num_int)

    if num_int > 0 {
      break
    } else {
      fmt.Println("Invlaid input. Enter number greater than 0:")
    }
  }
  return num_int
}

func build_slice(size int) []int {
  var intSclice []int
  for i := 0; i <= size; i++ {
    intSclice = append(intSclice, i)
  }
  return intSclice
}

func main () {
  //get int
  num := get_number()
  
  //build slice
  resultSlice := build_slice(num)

  //print slice test
  fmt.Println(resultSlice)
}
