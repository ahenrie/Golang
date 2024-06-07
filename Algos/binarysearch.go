package main

import (
  "fmt"
  //"strconv"
)
//we will get user input for a number then build a slice from 0 to the input 

//var int inputNumber
//fmt.Println("Enter a number for array size:")
//fmt.Scanln(&inputNumber)

func get_number() int {
  //var input string 
  var num int

  for {
    fmt.Println("Enter a number greater than 0:")
    fmt.Scanln(&num)

    //convet string to an int
    //num, err := strconv.Atoi(input)

    //check error and value of num greater than 0
    if num <= 0 {
      fmt.Println("Invalid input. Try again.")
    } else {
      break //exit if input is valid
    }
  }
  return num 
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
  sNum := get_number()
  
  fmt.Println(sNum)
  //build slice
  resultSlice := build_slice(sNum)

  //print slice test
  fmt.Println(resultSlice)
}
