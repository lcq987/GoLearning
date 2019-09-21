package main

import "fmt"

func main() {
	var x byte = 31
	var y byte = 27
	z := x &^ y
	z = z>>1
	fmt.Printf("%02d: %08b\n", x, x);
	fmt.Printf("%02d: %08b\n", y, y);
	fmt.Printf("%02d: %08b\n", z, z);
}
