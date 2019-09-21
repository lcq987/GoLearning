package lenconv

func MToF(m Meter) Foot {
  return Foot(m * 3.2808399)
}

func FToM(f Foot) Meter {
  return Meter(f * 0.3048)
}
