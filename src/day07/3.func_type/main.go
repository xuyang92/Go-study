package main

import "fmt"

// 函数类型与变量
type calculation func(int, int) int
func add(x, y int) int {
	return x + y
}

func f1() {
	fmt.Println("hello 沙河")
}

func f2() int {
	return 10
}

func f4(x, y int) int {
	return x + y
}

// 高阶函数
// 函数也可以作为参数的类型
func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

// 函数还可以作为返回值
func f5(x func() int) func(int, int) int {
	ret := func(a, b int) int {
		return a + b
	}
	return ret
}

func main() {
	var c calculation 				// 声明一个calculation类型的变量c
	c = add 						// 把add赋值给c
	fmt.Printf("type of c:%T\n", c) // type of c:main.calculation
	fmt.Println(c(1, 2)) 			// 像调用add一样调用c
	f := add 						// 将函数add赋值给变量f
	fmt.Printf("type of f:%T\n", f) // type of f:func(int, int) int
	fmt.Println(f(10, 20)) 			// 像调用add一样调用f }

	a:= f1
	fmt.Printf("%T\n", a)
	b := f2
	fmt.Printf("%T\n", b)

	f3(f2)
	f3(b)
	fmt.Printf("%T\n", f4)

	f7 := f5(f2)
	fmt.Printf("%T\n", f7)
}
