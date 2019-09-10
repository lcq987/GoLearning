package main

import (
    "fmt"
    "os"
)

func main()  {
  var s, sep string
  for _, arg := range os.Args[1:]{
    s += sep + arg
    sep = " "
  }
  fmt.Println(s)

  //修改echo程序，使其打印每个参数的索引和值，每个一行。
  for index, arg := range os.Args[:]{
    fmt.Println("index:", index, "parameter:", arg)
  }
}
