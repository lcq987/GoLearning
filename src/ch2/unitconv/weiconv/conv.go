package weiconv

func KToP(k Kilogram) Pound {
  return Pound(k * 2.2046226)
}

func PToK(p Pound) Kilogram {
  return Kilogram(p * 0.4535924)
}
