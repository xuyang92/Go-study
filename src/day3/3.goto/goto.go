package main

import "fmt"

func main() {
	var name string
	fmt.Println("请输入姓名：")
	fmt.Scanln(&name)
	if name == "svip" {
		goto SVIP
	} else if name == "vip" {
		goto VIP
	}

	fmt.Println("预约...")
VIP:
	fmt.Println("排号...")
SVIP:
	fmt.Println("进入...")
}
