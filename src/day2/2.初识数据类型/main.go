package main

import "fmt"

func main() {
	//整型
	fmt.Println(666)
	fmt.Println(6 + 9)
	fmt.Println(6 - 9)
	fmt.Println(6 * 9)
	fmt.Println(16 / 9)
	fmt.Println(16 % 9)

	//字符串，特点：通过双引号
	fmt.Println("曹煦阳")
	fmt.Println("真帅")
	fmt.Println("曹煦阳" + "真帅")
	//对比
	fmt.Println("1" + "2") 	//结果："12"
	fmt.Println(1 + 2) 		//结果：3

	//布尔类型，真假
	fmt.Println(1>2)	//false
	fmt.Println(1<2)	//true
	//超前，条件语句
	if 2 > 1 {
		fmt.Println("你请我吃饭")
	} else {
		fmt.Println("我请你吃饭")
	}
}
