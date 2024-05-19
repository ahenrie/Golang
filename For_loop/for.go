package main

import "fmt"

func main() {

  i := 1
  for i <= 3 {
    fmt.Println(i)
    i = i + 1
  }

  for j:=0; j < 3; j++ {
    fmt.Println(j)
  }

  //use a range for iteration of an int
  for i:= range 3 {
    fmt.Println("range",i)
  }
  
  //eternal loop without break
  for {
    fmt.Println("loop")
    break
  }
  
  //continue to next iteration
  for n:=range 6 {
    if n%2 == 0 {
      continue
    }
    fmt.Println(n)
  }

  fmt.Println("iterating over strings")
  //iterating over strings
  b := "Backwards"
  for i:=len(b) - 1; i >= 0; i-- {
    fmt.Print(string(b[i]))
  } 
  fmt.Println("")

  s := "Hello, World!"
  for _, char := range s {
    fmt.Print(string(char))
  }
}
