package main

import "fmt"

func main() {
	// 自动扩容
	v1 := make([]int, 1, 3)
	fmt.Println(len(v1), cap(v1))
	v2 := append(v1, 99)
	fmt.Println(len(v2), cap(v2))
	v3 := append(v2, 33)
	fmt.Println(len(v3), cap(v3))
	v1[0] = 123
	fmt.Println(v1, v2, v3)

	// 扩容前和扩容后的内存是不同的。
	m1 := []int{11, 22, 33}
	m2 := append(m1, 44)
	fmt.Println(len(m1), cap(m1))
	fmt.Println(len(m2), cap(m2))
	m1[0] = 321
	fmt.Println(m1, m2)

	// 切片
	p1 := []int{11, 22, 33, 44, 55, 66}
	p2 := p1[1:3]
	p3 := p1[1:]
	p4 := p1[:3]
	fmt.Println(p1, p2, p3, p4)

	// 追加
	n1 := []int{11, 22, 33}
	n2 := append(n1, 44)
	n3 := append(n1, 55, 66, 77, 88)
	n4 := append(n1, []int{100, 200, 300}...)
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)

	// 删除
	s1 := []int{11, 22, 33, 44, 55, 66}
	deleteIndex := 2
	result := append(s1[:deleteIndex], s1[deleteIndex+1:]...)
	fmt.Println(result)  	//[11 22 44 55 66]
	fmt.Println(s1)			//[11 22 44 55 66 66]

	// 插入
	r1 := []int{11, 22, 33, 44, 55, 66}
	insertIndex := 3 		//在索引3的位置插入99
	insertResult := make([]int, 0, len(r1)+1)
	insertResult = append(insertResult, r1[:insertIndex]...)
	insertResult = append(insertResult, 99)
	insertResult = append(insertResult, r1[insertIndex:]...)
	fmt.Println(r1, insertResult)

	// 循环
	q1 := []int{11, 22, 33, 44, 55, 66}
	for i:=0; i < len(q1); i++ {
		fmt.Println(i, q1[i])
	}
	for index, value := range q1 {
		fmt.Println(index, value)
	}

	// 切片嵌套
	ss1 := []int{11, 22, 33, 44, 55, 66}
	ss2 := [][]int{[]int{11, 22, 33}, []int{44, 55}}
	ss3 := [][2]int{[2]int{11, 22}, [2]int{33, 44}}
	fmt.Println(ss1)
	fmt.Println(ss2)
	fmt.Println(ss3)

	ss1[0] = 111111
	ss2[0][2] = 222222
	ss3[1][0] = 99999
	fmt.Println(ss1)
	fmt.Println(ss2)
	fmt.Println(ss3)
}
