package main

import "fmt"

func main() {
	r1 := 5 & 99
	r2 := 5 | 99
	r3 := 5 ^ 99
	r4 := 5 << 2
	r5 := 5 >> 1
	r6 := 5 &^ 99
	fmt.Println(r1, r2, r3, r4, r5, r6)
}
