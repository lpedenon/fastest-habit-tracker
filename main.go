package main

import (
  "fmt"
  "math"
  "math/rand"
  //"strconv"
) 

var c, python, java = true, false, "hello"

func main() {
  fmt.Println(rand.Int())
  fmt.Println("hello world")
  fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
  g := 3 // to declare use :=
  g = 4
  g += 4
  g++
  a, b := sum_product(3, 4)
  fmt.Println(a, b)
  var i int
  fmt.Println(i, c, python, java)
}

func sum_product(x, y int) (int, int) {
  return x + y, x * y
}
