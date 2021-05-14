package main

import "fmt"

func main() {
	//内置函数
	print("111111\n")
	print("222222\n")

	println("333333")
	println("444444")

	fmt.Print("我我我\n")
	fmt.Print("妮妮妮妮\n")
	fmt.Println("他他他他")
	fmt.Println("哈哈哈哈哈")

	//fmt包 扩展：格式化输出
	fmt.Printf("我你他%s，决定%d号一起踢足球，每个人需花费%f元场地费，最后花费%.2f元打车回家，获胜概率为%d%%，真是快乐的一天！\n",
		"组队", 20, 5.53, 10.339, 25)
}