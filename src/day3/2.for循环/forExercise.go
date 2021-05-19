package main

import "fmt"

func main() {
	//练习题1：猜数字，设定一个理想数字比如66，一直提示让用户输入数字，
	//如果比66大，则显示猜测的结果大了；如果比66小，则显示猜测的结果小了；
	//只有输入等于66，显示猜测结果正确，然后退出循环。
	/*
	data := 66
	flag := true
	for flag {
		var userInputNumber int
		fmt.Print("请输入数字：")
		fmt.Scanln(&userInputNumber)
		if userInputNumber < data {
			fmt.Println("小了")
		} else if userInputNumber > data {
			fmt.Println("大了")
		} else {
			fmt.Println("恭喜你猜对了！")
			flag = false 		//或者用break关键字退出循环
		}
	}
	*/

	//练习题2：使用循环输出1-100所有整数;
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}

	//练习题3：使用循环输出1、2、3、4、5、6、8、9、10，即：10以内除7以外的整数
	for i := 1; i <= 10; i++ {
		if i != 7 {
			fmt.Println(i)
		}
	}

	//练习题4：输出1-100内的所有奇数
	for i := 1; i <= 100; i++ {
		if i%2 == 1 {
			fmt.Println(i)
		}
	}

	//练习题5：输出1-100内的所有偶数
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	//练习题6：求1-100内的所有整数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	//练习题7：输出10-1的所有整数
	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}
}
