package main

import (
  "fmt"
  "math"
)

const s string = "constant"

func main() {
  fmt.Println(s)

  const n = 500000000
  
  //arbritrary precision 
  const d = 3e20 / n
  fmt.Println(d)

  //explicit conversion
  fmt.Println(int64(d))

  fmt.Println(math.Sin(n))
}
