package main

import "fmt"

func main() {
	// var a int
	// a = 4
	a := 4
	b := 2

	fmt.Printf("a&b = %v\n", a&b)
	fmt.Printf("a|b = %v\n", a|b)
	fmt.Printf("a^b = %v\n", a^b)
	fmt.Printf("a >> 1 = %v\n", a<<1)
	fmt.Printf("a << 1 = %v\n", a>>1)
}
