package popcount

var pc [256]byte  //2的8次方为256，将所有的数按8bit为单位分割，分割出来的每一段最大为256，
                  //因此作长度为256的表记录0-255中每个数的二进制中含1的个数
func init() {
  for i, _ := range pc {
    pc[i] = pc[i/2] + byte(i&1)
  }
}

func Popcount(x uint64) int {
  return int(pc[byte(x>>(0*8))]+  //byte == uint8，byte(x)输出的是x低8位的值
        pc[byte(x>>(1*8))] +
        pc[byte(x>>(2*8))] +
        pc[byte(x>>(3*8))] +
        pc[byte(x>>(4*8))] +
        pc[byte(x>>(5*8))] +
        pc[byte(x>>(6*8))] +
        pc[byte(x>>(7*8))])
}
