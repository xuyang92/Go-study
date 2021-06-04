package main

import "fmt"

// 闭包

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

// 定义一个函数对f2进行包装
func f3(f func(int, int), m, n int) func() {
	tmp := func() {
		f(m, n)
	}
	return tmp
}

//func f3(x, y int) func() {
//	tmp := func() {
//		fmt.Println(x+y)
//	}
//	return tmp
//}

func main() {
	// 把原来需要传递两个int类型的参数包装成一个不需要传参的参数
	ret := f3(f2, 100, 200)
	f1(ret)
	//ret := f3(100, 200)
	//f1(ret)
}
