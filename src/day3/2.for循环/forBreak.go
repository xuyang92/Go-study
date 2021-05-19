package main

import "fmt"

func main() {
	//练习题1：猜数字，设定一个理想数字比如66，一直提示让用户输入数字，
	//如果比66大，则显示猜测的结果大了；如果比66小，则显示猜测的结果小了；
	//只有输入等于66，显示猜测结果正确，然后退出循环。

	fmt.Println("开始")
	data := 66
	for {
		var userInputNumber int
		fmt.Print("请输入数字：")
		fmt.Scanln(&userInputNumber)
		if userInputNumber < data {
			fmt.Println("小了")
		} else if userInputNumber > data {
			fmt.Println("大了")
		} else {
			fmt.Println("恭喜你猜对了！")
			break
		}
	}
	fmt.Println("结束")


	for i := 1; i < 3; i++ {
		for j := 1; j < 5; j++ {
			if j == 3 {
				break
			}
			fmt.Println(i, j)
		}
	}
}
