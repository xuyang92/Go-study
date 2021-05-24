package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
	n1 := 666
	n2 := 9223372036854775807	//n2在加上1或者n3在减去1，编译运行就报错。
	n3 := -9223372036854775808	//数值超出整型在64位操作系统上表示的范围。
	fmt.Println(n1, n2, n3)

	var v1 int8 = 2
	var v2 int16 = 6
	v3 := int16(v1) + v2
	fmt.Println(v3)

	var v4 int16 = 128
	v5 := int8(v4)
	fmt.Println(v5)


	//整型转换为字符串
	s1 := 199
	result1 := strconv.Itoa(s1)
	fmt.Println(result1, reflect.TypeOf(result1))

	//字符串转换为整型：转换可能失败，返回err
	s2 := "666"
	result2, err := strconv.Atoi(s2)
	if err == nil {
		fmt.Println("转换成功", result2)
	} else {
		fmt.Println("转换失败")
	}

	//整型转换为其他进制
	s3 := 5
	r1 := strconv.FormatInt(int64(s3), 2)
	fmt.Println(r1, reflect.TypeOf(r1))

	//其他进制转换为整型
	//通过ParseInt将字符串转换为整型时，本质上与Atoi是相同的。
	s4 := "1000111100101"
	//s4要转换的文本，2把文本当作二进制转换成整型，16转换的过程中对结果进行约束。
	r2, err2 := strconv.ParseInt(s4, 2, 16)
	if err2 == nil {
		fmt.Println("转换成功", r2, reflect.TypeOf(r2))
	} else {
		fmt.Println("转换失败")
	}
	*/

	//将十进制14转换成16进制的字符串
	v1 := strconv.FormatInt(14, 16)
	fmt.Println(v1)

	//将二进制"10011"转换成十进制的整型
	v2, _ := strconv.ParseInt("10011", 2, 0)
	fmt.Println(v2)

	//将二进制"1001101"转换成十六进制的字符串
	v3, _ := strconv.ParseInt("1001101", 2, 0)
	v4 := strconv.FormatInt(v3, 16)
	fmt.Println(v4)
}
