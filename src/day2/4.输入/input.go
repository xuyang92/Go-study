package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//示例1：fmt.Scan
	//特别说明：fmt.Scan要求输入两个值，必须输入两个，否则会一直等待。
	/*
	var name string
	var age int
	fmt.Print("请输入用户名：")
	//当使用Scan时，会提示用户输入
	//用户输入完成之后，会得到两个值：
	//1、count，用户输入了几个值；
	//2、err，输入错误信息
	_, err := fmt.Scan(&name, &age)
	if err == nil {
		fmt.Println(name, age)
	} else {
		fmt.Println("用户输入数据错误", err)
	}
	*/

	//示例2：fmt.Scanln
	//特别说明：fmt.Scanln要求输入两个值，不一定输入两个值，等待回车就会结束输入。
	/*
	var name string
	var age int
	fmt.Print("请输入用户名：")
	//当使用Scan时，会提示用户输入
	//用户输入完成之后，会得到两个值：
	//1、count，用户输入了几个值；
	//2、err，输入错误信息
	_, err := fmt.Scanln(&name, &age)
	if err == nil {
		fmt.Println(name, age)
	} else {
		fmt.Println("用户输入数据错误", err)
	}
	*/

	//示例3：fmt.Scanf  格式化输入
	/*
	var name string
	var age int
	fmt.Print("请输入用户名：")
	count, err := fmt.Scanf("name%s age%d", &name, &age)
	if err == nil {
		fmt.Println("输入成功", count)
		fmt.Println(name, age)
	} else {
		fmt.Println("输入失败", count, err)
	}
	*/

	//当输入例如："盼望着盼望着 东风来了"
	//message的值为空格的前半部分："盼望着盼望着"
	//var message string
	//fmt.Print("请输入个人信息：")
	//fmt.Scanln(&message)
	//fmt.Println(message)

	reader := bufio.NewReader(os.Stdin)
	// line从stdin中读取一行数据（字节集合--->转化成字符串）
	// reader默认一次最多能读4096个字节（4096/3个汉字）
	// 1、一次性读完，isPrefix = false
	// 2、先读一部分，isPrefix = true（输入内容大于4096个字节）
	line, isPrefix, err := reader.ReadLine()
	data := string(line)
	fmt.Println(data, isPrefix, err)

}
