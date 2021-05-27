package main

import "fmt"

func main() {
	var name string = "曹煦阳"
	//1.索引获取字节
	v1 := name[0]
	fmt.Println(v1)	//230

	//2.切片获取字节区间
	v2 := name[0:3]
	fmt.Println(v2)

	//3.手动循环获取所有字节
	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
	}

	//4.for range循环获取所有字符
	for index, item := range name {
		fmt.Println(index, item)
	}

	//5.转换成rune集合
	datalist := []rune(name)
	fmt.Println(datalist[0], string(datalist[0]))
}
