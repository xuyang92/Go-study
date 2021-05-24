package main

import (
	"fmt"
	"strconv"
)

func main() {
	//定义字符串，字符串是以什么形式存在于Go编译器（Utf-8编码）
	name := "曹煦阳"
	// 其他语言：name[0] 曹；name[1] 煦;name[2] 阳
	//曹   Go语言下标取到的是每一个字节。
	fmt.Println(name[0], strconv.FormatInt(int64(name[0]), 2))	//230
	fmt.Println(name[1], strconv.FormatInt(int64(name[1]), 2))	//155
	fmt.Println(name[2], strconv.FormatInt(int64(name[2]), 2))	//185

	//煦
	fmt.Println(name[3])
	fmt.Println(name[4])
	fmt.Println(name[5])

	//阳
	fmt.Println(name[6])
	fmt.Println(name[7])
	fmt.Println(name[8])
}
