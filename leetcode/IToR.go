package main

import "fmt"

func main() {
  fmt.Println(intToRoman(3994))
}

func intToRoman(num int) string {
    var res string
    var roman = map[int]string{
        1000: "M",
        900 : "CM",
        500 : "D",
        400 : "CD",
        100 : "C",
        90  : "XC",
        50  : "L",
        40  : "XL",
        10  : "X",
        9   : "IX",
        5   : "V",
        4   : "IV",
        1   : "I",
    }
    ara := [...]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    for num > 0{
        for _, v := range ara {
            if num >= v{
              fmt.Println(v)
              res += roman[v]
              num -= v
              break
            }
        }
    }
    return res
}
