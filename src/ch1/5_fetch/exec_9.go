//练习 1.9：修改fetch打印出HTTP协议的状态码，可以从resp.Status变量得到该状态码。


package main

import(
  "fmt"
  "os"
  "net/http"
  "strings"
)

func main() {
  for _, url := range os.Args[1:]{
    if strings.HasPrefix(url, "http://") == false {
      url = "http://" + url
    }
    resp, err := http.Get(url)
    if err != nil {
      fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
      os.Exit(1)
    }
    fmt.Println("http status: ", resp.Status)
  }
}
