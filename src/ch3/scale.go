package main

import (
  "fmt"
  "math"
  "strconv"
  //"os"
)

func main()  {
  fmt.Println(NToN("01101111", 2, 7))
}

func NToN(a string, n int, m int) string {
  dec := NToD(a, n)
  aim := DToN(dec, m)
  return aim
}

func NToD(num string, n int) int{
  var dec int = 0
  var count float64 = 0
  mod := map[string]int{
    "A": 10, "B": 11, "C": 12, "D": 13, "E": 14, "F": 15,
  }
  for i := len(num); i > 0; i--{
    v := num[i-1:i]
    if(n == 16){
      if _, ok := mod[v]; ok{
        dec += int(math.Pow(16, count))*mod[v]
      }else{
        s, _ := strconv.Atoi(v)
        dec += int(math.Pow(16, count))*s
      }
    }else{
      s, _ := strconv.Atoi(v)
      dec += int(math.Pow(float64(n), count))*s
    }
    count += 1
  }
  return dec
}

func DToN(num int, n int) string {
  var mode []string = []string{"A", "B", "C", "D", "E", "F"};
  var result string
  if n == 16 {
    for num > 0 {
      res := num % 16
      num = num / 16
      if res >= 10{
        result = mode[res-10]+result
      }else{
        result = strconv.Itoa(res)+result
      }
    }
  }else{
    for num > 0 {
      res :=  num % n
      num = num / n
      result = strconv.Itoa(res)+result
    }
  }
  return result
}
