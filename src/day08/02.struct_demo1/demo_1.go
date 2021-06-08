package main

import "fmt"

// 结构体

type person struct {
	name string
	age int
	gender string
	hobby []string
}

func main() {
	// 声明一个person类型的变量p
	var p person
	// 通过字段赋值
	p.name = "cxy"
	p.age = 29
	p.gender = "男"
	p.hobby = []string{"足球", "跑步", "双色球", "游泳"}
	fmt.Println(p)
	fmt.Printf("%T\n", p)
	// 访问变量p的字段
	fmt.Println(p.name)

	// 匿名结构体：多用于临时场景
	var s struct {
		x string
		y int
	}
	s.x = "fffff"
	s.y = 100
	fmt.Printf("type:%T value:%v\n", s, s)
}
