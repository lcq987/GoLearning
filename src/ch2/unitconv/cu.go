package main

import(
  "fmt"
  "os"
  "strconv"
  "./lenconv"
  "./tempconv"
  "./weiconv"
)

func main() {
  for _, arg := range os.Args[1:]{
    t, err := strconv.ParseFloat(arg, 64)
    if err != nil {
      fmt.Fprintf(os.Stderr, "cf: %v\n", err)
      os.Exit(1)
    }
    m := lenconv.Meter(t)
    ft := lenconv.Foot(t)
    c := tempconv.Celsius(t)
    f := tempconv.Fahrenheit(t)
    kg := weiconv.Kilogram(t)
    lb := weiconv.Pound(t)
    fmt.Printf("%s = %s, %s = %s, %s = %s, %s = %s, %s = %s, %s = %s\n", m, lenconv.MToF(m), ft, lenconv.FToM(ft), c, tempconv.CToF(c), f, tempconv.FToC(f), kg, weiconv.KToP(kg), lb, weiconv.PToK(lb))
  }
}
