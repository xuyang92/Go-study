package main

import (
	"fmt"
	"strings"
)

// 闭包是什么？
// 闭包是一个函数，这个函数包含了他外部作用域的一个变量。

// 闭包的底层原理：
// 1.函数可以作为返回值；
// 2.函数内部查找变量的顺序，先在自己内部找，找不到往外层找。

// adder函数返回值为func(int) int类型的函数
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	// 示例
	var f = adder()		// f为adder函数的返回值，即func(int) int函数
	fmt.Println(f(10)) 	//10
	fmt.Println(f(20)) 	//30
	fmt.Println(f(30)) 	//60
	f1 := adder()
	fmt.Println(f1(40)) //40
	fmt.Println(f1(50)) //90

	//进阶示例1
	var f_2 = adder2(10)
	fmt.Println(f_2(10)) //20
	fmt.Println(f_2(20)) //40
	fmt.Println(f_2(30)) //70
	f1_2 := adder2(20)
	fmt.Println(f1_2(40)) //60
	fmt.Println(f1_2(50)) //110

	//进阶示例2
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt

	//进阶示例3
	f1_3, f2_3 := calc(10)
	fmt.Println(f1_3(1), f2_3(2)) //11 9
	fmt.Println(f1_3(3), f2_3(4)) //12 8
	fmt.Println(f1_3(5), f2_3(6)) //13 7
}



