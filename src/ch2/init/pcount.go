package main

import (
  "fmt"
  "./popcount"
)

func main() {
  var num uint64 = 112433
  count := popcount.Popcount(num)
  fmt.Println(num, "中的含二进制1bit的个数为", count)
}
