package main

import "fmt"

func main() {
	v1 := 1
	v2 := v1
	fmt.Printf("v1的内存地址：%p \n", &v1)	//0xc0000180a8
	fmt.Printf("v2的内存地址：%p \n", &v2)	//0xc0000180c0

	b1 := false
	b2 := b1
	fmt.Printf("b1的内存地址：%p \n", &b1)	//0xc0000180c8
	fmt.Printf("b2的内存地址：%p \n", &b2)	//0xc0000180c9

	f1 := 3.14
	f2 := f1
	fmt.Printf("f1的内存地址：%p \n", &f1)	//0xc0000180e0
	fmt.Printf("f2的内存地址：%p \n", &f2)	//0xc0000180e8

	s1 := "曹煦阳"
	s2 := s1
	fmt.Printf("s1的内存地址：%p \n", &s1)	//0xc000042240
	fmt.Printf("s2的内存地址：%p \n", &s2)	//0xc000042250

	arr1 := [2]int{6, 9}
	arr2 := arr1
	fmt.Printf("arr1的内存地址：%p \n", &arr1)	//0xc0000180f0
	fmt.Printf("arr2的内存地址：%p \n", &arr2)	//0xc000018100
	arr1[0] = 1111
	fmt.Println(arr1, arr2)

	sli1 := []int{6, 9}
	sli2 := sli1
	fmt.Printf("sli1的内存地址：%p \n", &sli1)	//0xc000004078
	fmt.Printf("sli2的内存地址：%p \n", &sli2)	//0xc000004090
	sli1[0] = 1111
	sli1 = append(sli1, 999)
	fmt.Println(sli1, sli2)
}
