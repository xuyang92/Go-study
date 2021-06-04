package main

import "fmt"

// defer

// defer多用于函数结束之前释放资源（文件句柄、数据库连接、socket连接）
func deferDemo()  {
	fmt.Println("start")
	// defer把它后面的语句延迟到函数即将返回的时候再执行;
	defer fmt.Println("嘿嘿嘿")
	defer fmt.Println("呵呵呵")
	defer fmt.Println("哈哈哈")
	fmt.Println("end")
}

func main() {
	deferDemo()
}
