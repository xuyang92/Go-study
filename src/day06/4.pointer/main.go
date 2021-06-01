package main

import "fmt"

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

func main() {
	// 1.&:取地址
	n := 18
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p)
	// 2.*:根据地址取值
	m := *p
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	a := 10
	modify1(a)
	fmt.Println(a)
	modify2(&a)
	fmt.Println(a)
}
