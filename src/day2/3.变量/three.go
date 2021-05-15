package main

import "fmt"

//全局变量(不能以省略的方式简写)
var school string = "w3cschool"	//可以
//var school = "w3cschool" 		//可以
//school := "w3cschool" 		//不可以，ERROR

var (
	global_v1 = 123
	global_v2 = "你好"
	global_v3 int
)

func main() {
	//声明+赋值
	//var sd string = "曹老汉真帅"
	//fmt.Println(sd)
	//
	//var age int = 29
	//fmt.Println(age)
	//
	//var flag bool = true
	//fmt.Println(flag)

	//先声明后赋值
	//var sd string
	//sd = "曹老汉真帅"
	//fmt.Println(sd)

	//编写代码省事
	//文本，请将文本输出3次：“伤情最是晚凉天，憔悴斯人不堪怜”
	//var message string = "伤情最是晚凉天，憔悴斯人不堪怜"
	//fmt.Println(message)
	//fmt.Println(message)
	//fmt.Println(message)
	//for i:=1; i<=3;i++ {
	//	fmt.Println(message)
	//}

	//存储数据
	//存储结果，方便之后使用;
	//var v1 int = 99
	//var v2 int = 33
	//var v3 int = v1 + v2
	//var v4 int = v1 * v2
	//var v5 int = v1 + v2 + v3 + v4
	//fmt.Println(v1, v2, v3, v4, v5)
	//存储输入的值
	//var name string
	//fmt.Scanf("%s", &name)   	//用户输入字符串并赋值给name变量
	////fmt.Println(name)
	//if  name == "曹煦阳" {
	//	fmt.Println("用户名输入正确")
	//} else {
	//	fmt.Println("用户名输入失败")
	//}
	//变量名要求
	//1、变量名必须只包含：字母、数字、下划线
	//2、数字不能开头
	//3、不能使用go语言内置的关键字（break、default、func、interface、select、case、defer、go、map、 struct、chan、
	//else、goto、package、switch、const、fallthrough、if、range、type、continue、for、inport、return、var共25个关键字）
	//1、变量名见名知意；2、驼峰式命名：例如myBossName

	//变量简写
	//声明+赋值
	//var myName string = "曹煦阳"
	//var myName = "曹煦阳"  		//简写方式1
	//myName := "曹煦阳"				//简写方式2，推荐

	//var myName, myAge, myHeight string = "曹煦阳", "29岁", "65kg"
	//var myName, myAge, myHeight = "曹煦阳", "29岁", "65kg"
	//myName, myAge, myHeight := "曹煦阳", "29岁", "65kg"
	//fmt.Println(myName, myAge, myHeight)

	//先声明再赋值
	//var myName, myAge, myHeight string
	//myName = "曹煦阳"
	//myAge = "29岁"
	//myHeight = "65kg"
	//fmt.Println(myName, myAge, myHeight)

	//因式分解，例如：声明5个变量，分别有字符串、整型
	/*var (
		myName = "曹煦阳"
		myAge = 29
		myHeight = "65kg"
		myHabbit = "marathon"
		mySalary string				//只声明但不赋值，默认值""
		myLength int				//只声明但不赋值，默认值0
		hasWife bool				//只声明但不赋值，默认值false
	)
	fmt.Println(myName, myAge, myHeight, myHabbit, mySalary, myLength, hasWife)*/

	///////扩展------变量声明不使用go编译器会报错-------///////

	//变量作用域
	//如果我们定义了大括号，那么在大括号中定义的变量
	//***不能被上级调用
	//***可以在同级使用
	//***可以在子级使用
	name_1 := "曹煦阳"			//局部变量
	fmt.Println(name_1)
	if true {
		age_1 := 29				//局部变量
		fmt.Println(age_1)
		name_1 := "cxy"			//局部变量
		fmt.Println(name_1)
	}
	fmt.Println(name_1)

	//全局变量和局部变量
	//**1、全局变量，未写在函数中的变量称为全局变量，
	//****不可以使用v1 := xx的方式进行简写，可以基
	//****于因式分解的方式声明多个变量；项目中寻找变
	//****量时的最后一环。
	//**2、局部变量，编写在作用域{}里面的变量，可以
	//****使用任意方式简写，可以基于因式分解的方式
	//****声明多个变量。
	fmt.Println(school)
	fmt.Println(global_v1, global_v2, global_v3)

	//-------赋值及内存相关---------//
	//*****注意事项****
	//**使用int、string、bool这三种数据类型时，如果遇到变量的赋值则会拷贝一份。【值类型】
	name_ptr := "cxy"
	nick_name_ptr := name_ptr
	fmt.Println(name_ptr, &name_ptr)
	fmt.Println(nick_name_ptr, &nick_name_ptr)
	name_ptr = "曹煦阳"
	fmt.Println(name_ptr, &name_ptr)
	fmt.Println(nick_name_ptr, &nick_name_ptr)

	//-------常量---------//
	//const age int = 98
	//const age = 98
	//fmt.Println(age)
	//常量因式分解
	//const (
	//	v1 = 123
	//	v2 = 456
	//)
	//fmt.Println(v1, v2)

	const (   //可有可无，当作一个声明常量时的计数器（值为0）
		v1 = iota
		v2
		v3
		v4
		v5
	)
	fmt.Println(v1,v2,v3,v4,v5)
}
