package main

import (
	"fmt"
	"math/big"
)

func main() {
	//第一步：创建一个超大整型的一个对象
	var v1 big.Int

	//第二步：在超大整型对象中写入值
	v1.SetInt64(1990)
	fmt.Println(v1.String())
	//推荐下面这种写法
	v1.SetString("901231231200000000000", 10)
	fmt.Println(v1.String())

	//var v2 *big.Int   //一般用不着（直接复制的时候使用）
	v3 := new(big.Int)
	v3.SetInt64(1990)
	fmt.Println(v3)
	//推荐下面这种写法
	v3.SetString("901231231200000000000", 10)
	fmt.Println(v3)

	n1 := new(big.Int)
	n1.SetString("12330025546886220014856222", 10)
	//n1.SetInt64(89)
	//n1 := big.NewInt(89)
	n2 := new(big.Int)
	n2.SetString("52244885214524488522233", 10)
	//n2.SetInt64(99)
	//n2 := big.NewInt(99)
	result := new(big.Int)
	result.Add(n1, n2)
	fmt.Println(result.String())

	var r1 big.Int
	r1.SetString("92222000365477582222255662182220000", 10)
	var r2 big.Int
	r2.SetString("2", 10)
	//result_1 := new(big.Int)
	var result_1 big.Int
	result_1.Add(&r1, &r2)
	fmt.Println(result_1.String())
}
