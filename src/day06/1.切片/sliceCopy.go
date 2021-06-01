package main

import "fmt"

func main() {
	// copy()赋值切片
	a1 := []int{1, 3, 5}
	a2 := a1
	var a3 = make([]int, 3, 3)
	copy(a3, a1)
	fmt.Println(a1, a2, a3)
	a1[0] = 100
	fmt.Println(a1, a2, a3)

	//由于切片是引用类型，所以a1和a2其实都指向了同一块内存地址，修改a1的同时a2的值也会发生变化。
	//Go语言内建的copy()函数可以迅速地将一个切片的数据复制到另外一个切片空间中。
}
