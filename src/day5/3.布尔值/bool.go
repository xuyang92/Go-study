package main

import (
	"fmt"
	"strconv"
)

func main() {
	//字符串转换布尔类型
	//true:"1", "t", "T", "true", "TRUE", "True"
	//false:"0", "f", "F", "false", "FALSE", "False"
	result, err := strconv.ParseBool("true")
	fmt.Println(result, err)

	//布尔类型转换字符串
	v2 := strconv.FormatBool(true)
	fmt.Println(v2)
}
