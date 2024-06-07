package main

import "fmt"

func main() {
  //create an array that will hold 5 ints
  var a [5]int
  fmt.Println("emp:", a)

  //all indices will be written with a 0

  //access 5th int 
  a[4] = 100
  fmt.Println("set:", a)
  fmt.Println("get:", a[4])

  //print length of an array
  fmt.Println("len:", len(a))

  //declare and initialize in one line
  b := [5]int{1, 2, 3, 4}

  //same but use the compiler to count
  b = [...]int{0, 0:100}
  fmt.Println(b)


}
