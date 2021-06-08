package main

import "fmt"

// 结构体是值类型

type person struct {
	name,gender string
}

// Go语言中函数参数永远是拷贝
func f(x person) {
	x.gender = "女"	//修改的是副本的gender
}

func f2(x *person) {
	// (*x).gender = "女"
	x.gender = "女"	// 语法糖，自动根据指针找对应的变量
}

func main() {
	var p person
	p.name = "cxy"
	p.gender = "男"
	f(p)
	fmt.Println(p.gender)	// 男
	f2(&p)
	fmt.Println(p.gender)	// 女

	// 结构体指针1
	var p2 = new(person)
	p2.name = "理想"
	fmt.Printf("%T\n", p2)
	fmt.Printf("%p\n", p2)

	// 结构体指针2
	// 使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
	p3 := &person{}
	fmt.Printf("%T\n", p3)     //*main.person
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", gender:""}
	p3.name = "小丽"
	p3.gender = "女"
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"小丽", gender:"女"}

	var p4 = person{
		name: "元帅",
		gender: "男",
	}
	fmt.Printf("%#v\n", p4)

	var p5 = &person{
		/*name:*/ "孙悟空",
		/*gender:*/ "男",
	}
	fmt.Printf("%#v\n", p5)
}
