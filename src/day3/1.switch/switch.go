package main

import "fmt"

func main() {
	var age int
	fmt.Scanln(&age)
	switch age {
	case 1:
		fmt.Println("等于1")
		fallthrough
	case 2:
		fmt.Println("等于2")
	case 3:
		fmt.Println("等于3")
		fallthrough
	case 4, 5, 6:
		fmt.Println("等于4或5或6")
	default:
		fmt.Println("都不满足")
	}

	//注意事项：数据类型一致的情况，正确：1>2、3+4，错误1>"3"、5+"6"
	//1、左花括号{必须与switch处于同一行
	//2、条件表达式不限制为常量或者整数
	//3、单个case中，可以出现多个结果选项
	//4、与C语言规则相反，Go语言不需要用break来明确退出一个case
	//5、只有在case中明确添加fallthrough关键字，才会继续执行紧跟的下一个case

}
