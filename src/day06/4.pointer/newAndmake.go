package main

import "fmt"

func main() {
	var a1 *int
	fmt.Println(a1)
	var a2 = new(int)
	fmt.Println(a2)
	fmt.Println(*a2)
	*a2 = 100
	fmt.Println(*a2)

}
