package main

import "fmt"

func main() {
	var name, address, action string
	fmt.Println("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Println("请输入地址：")
	fmt.Scanln(&address)
	fmt.Println("请输入动作：")
	fmt.Scanln(&action)

	result := fmt.Sprintf("我叫%s，在%s正在%s", name, address, action)
	fmt.Println(result)
}
