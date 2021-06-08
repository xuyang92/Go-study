package main

import "fmt"

// 递归：函数自己调用自己
// 递归适合处理那种问题相同、问题的规模越来越小的场景
// 递归一定要有一个明确的退出条件


// 阶乘：以下实例通过Go语言的递归函数实现阶乘；
func Factorial(n uint64)(result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

// 上台阶的面试题:
// n个台阶，一次可以走1步，也可以走2步，有多少种走法？
func stepCalc(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return stepCalc(n-1) + stepCalc(n-2)
}

func main() {
	var i int = 15
	fmt.Printf("%d的阶乘是%d\n", i, Factorial(uint64(i)))

	ret := stepCalc(4)
	fmt.Println(ret)
}
