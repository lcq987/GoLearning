package main

import "fmt"

func main() {
  var nums []int = []int{2, 7, 11, 15}
  target := 9
  fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
  hash := make(map[int]int)
  var r1, r2 int
  for k, v := range nums{
    r := target - v
    if i, ok := hash[r]; !ok{
      hash[v] = k
    }else{
      r1, r2 = i, k
    }
  }
  return []int{r1, r2}
}
