package weiconv

import "fmt"

type Kilogram float64
type Pound float64

func (kg Kilogram)String() string{
  return fmt.Sprintf("%gkg", kg)
}

func (p Pound)String() string{
  return fmt.Sprintf("%glb", p)
}
