package main

import "fmt"

func main() {
	//练习题3：使用循环输出1、2、3、4、5、6、8、9、10，即：10以内除7以外的整数
	for i := 1; i <= 10; i++ {
		if i == 7 {
			continue
		}
		fmt.Println(i)
	}
	
	//for循环嵌套
	for i := 1; i < 3; i++ {
		for j := 1; j < 5; j++ {
			fmt.Println(i, j)
		}
	}

	//for循环嵌套 + continue
	for i := 1; i < 3; i++ {
		for j := 1; j < 5; j++ {
			if j == 3 {
				continue
			}
			fmt.Println(i, j)
		}
	}
}
