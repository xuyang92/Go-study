package main

import "fmt"

//条件语句
func main() {
	//if true {
	//	fmt.Println("666")
	//} else {
	//	fmt.Println("999")
	//}
	//
	//if 1 > 2 {
	//	fmt.Println("666")
	//} else {
	//	fmt.Println("999")
	//}
	//
	//flag := true
	//if flag {
	//	fmt.Println("条件成立")
	//} else {
	//	fmt.Println("条件不成立")
	//}

	//练习题1：用户输入姓名，判断是否正确
	/*
	var name string
	fmt.Println("请输入姓名：")
	fmt.Scanln(&name)
	if name == "cxy" {
		fmt.Println("用户名输入正确")
	} else {
		fmt.Println("用户名输入错误")
	}
	*/

	//练习题2：用户输入数字，判断奇数还是偶数
	/*
	var number int
	fmt.Println("请输入数字：")
	fmt.Scanln(&number)
	if number % 2 == 0 {
		fmt.Println("您输入的是偶数")
	} else {
		fmt.Println("您输入的是奇数")
	}
	*/

	//练习题3：用户输入用户名和密码，判断是否正确
	/*
	var username, password string
	fmt.Println("请输入用户名：")
	fmt.Scanln(&username)
	fmt.Println("请输入密码：")
	fmt.Scanln(&password)
	if username == "cxy" && password == "cxy" {
		fmt.Println("欢迎登录gitHub")
	} else {
		fmt.Println("用户名或密码错误")
	}
	*/

	//练习题3：请输入用户名校验是否是VIP
	/*
	var username string
	fmt.Println("请输入用户名：")
	fmt.Scanln(&username)
	if username == "cxy" || username == "vip" {
		fmt.Println("vip用户")
	} else {
		fmt.Println("非vip用户")
	}
	*/

	//多条件语句，示例如下：
	/*
	var weight int
	fmt.Println("请输入你的体重（kg）：")
	fmt.Scanln(&weight)
	if weight < 50 {
		fmt.Println("体型瘦弱")
	} else if weight < 60 {
		fmt.Println("体型完美")
	} else if weight < 70 {
		fmt.Println("体型健康")
	} else {
		fmt.Println("体型臃肿")
	}
	*/

	//嵌套
	fmt.Println("欢迎致电10086，1：话费相关；2：业务办理；3：人工服务。")
	var number int
	fmt.Scanln(&number)
	if number == 1 {
		fmt.Println("话费服务，1：余额查询；2：话费缴纳。")
		var n1 int
		fmt.Scanln(&n1)
		if n1 == 1 {
			fmt.Println("余额查询")
		} else if n1 == 2 {
			fmt.Println("话费缴纳")
		} else {
			fmt.Println("输入错误")
		}
	} else if number == 2 {
		fmt.Println("业务办理")
	} else if number == 3 {
		fmt.Println("人工服务")
	} else {
		fmt.Println("输入错误")
	}
}
