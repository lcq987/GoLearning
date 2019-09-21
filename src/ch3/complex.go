package main

import (
  "fmt"
  "math/cmplx"
)

func main() {
  x := 3i
  y := 5 - 2i
  fmt.Println(x*y)
  fmt.Println(real(x*y))
  fmt.Println(imag(x*y))
  fmt.Println(cmplx.Sqrt(-2))
}
