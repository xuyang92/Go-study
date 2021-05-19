package main

import "fmt"

func main() {
	//示例1：无限循环
	/*
	fmt.Println("开始")
	for {
		fmt.Println("红鲤鱼与绿鲤鱼与驴")
		time.Sleep(time.Second * 1)		//等一秒在继续执行
	}
	fmt.Println("结束")
	*/

	//示例2：
	/*
	fmt.Println("开始")
	for 2 > 1 {
		fmt.Println("红鲤鱼与绿鲤鱼与驴")
		time.Sleep(time.Second * 1)		//等一秒在继续执行
	}
	fmt.Println("结束")
	*/

	//示例3：
	fmt.Println("开始")
	for number := 1; number < 5; number += 1 {
		fmt.Println("钓鱼要钓刀鱼，刀鱼要到岛上钓")
	}
	fmt.Println("结束")


}
