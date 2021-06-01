package main

import "fmt"

func main() {
	/*
	//方式一:先声明再赋值
	//内存中已开辟空间，内存初始化的值都是0
	var numbers [3]int
	numbers[0] = 333
	numbers[1] = 666
	numbers[2] = 999
	//方式二：声明+赋值
	var names = [3]string{"曹煦阳", "Bob", "John"}
	fmt.Println(names)
	//方式三：声明+赋值+指定位置
	var ages = [3]int{0:65, 2:16}
	fmt.Println(ages)
	//方式四：省略个数
	var names_1 = [...]string{"曹煦阳", "Bob", "John"}
	var ages_1 = [...]int{0:65, 2:16}
	fmt.Println(names_1, ages_1)

	//数组指针
	//声明：指针类型的数组{本质上是一个指针类型}
	//不会开辟内存来初始化数组中的值，numbers == nil
	var numbers_1 *[3]int
	//声明数组并初始化，返回的是指针类型
	numbers_2 := new([3]int)

	numbers_1 = &ages
	numbers_2 = &ages
	fmt.Println(*numbers_1, *numbers_2)
	*/

	nums := [3]int8{11, 22, 33}
	// %p表示格式化输出内存地址
	fmt.Printf("数组的内存地址：%p \n", &nums)
	fmt.Printf("数组第一个元素的内存地址：%p \n", &nums[0])
	fmt.Printf("数组第二个元素的内存地址：%p \n", &nums[1])
	fmt.Printf("数组第三个元素的内存地址：%p \n", &nums[2])

	nums32 := [3]int32{11, 22, 33}
	// %p表示格式化输出内存地址
	fmt.Printf("数组的内存地址：%p \n", &nums32)
	fmt.Printf("数组第一个元素的内存地址：%p \n", &nums32[0])
	fmt.Printf("数组第二个元素的内存地址：%p \n", &nums32[1])
	fmt.Printf("数组第三个元素的内存地址：%p \n", &nums32[2])

	names := [2]string{"曹煦阳", "cxy"}
	// %p表示格式化输出内存地址
	fmt.Printf("数组的内存地址：%p \n", &names)
	fmt.Printf("数组第一个元素的内存地址：%p \n", &names[0])
	fmt.Printf("数组第二个元素的内存地址：%p \n", &names[1])

	//拷贝
	names1 := [3]string{"曹煦阳", "Bob", "John"}
	names2 := names1
	names1[0] = "cxy"
	fmt.Println(names1, names2)

	// 1.长度
	fmt.Println(len(names))
	// 2.索引
	namesIndex := [2]string{"曹煦阳", "cxy"}
	data := namesIndex[0]
	fmt.Println(data)
	// 3.切片
	nums_slice := [3]int{11, 22, 33}
	numsSlice := nums_slice[0:2] //获取下标大于等于0及小于2的元素组成新的数组，即切片
	fmt.Println(numsSlice)
	// 4.循环
	nums_for := [3]int32{11, 22, 33}
	for i := 0; i < len(nums_for); i++ {	//手动循环
		fmt.Println(i, nums_for[i])
	}
	for key, item := range nums_for {		//for range循环
		fmt.Println(key, item)
	}

	// 数组嵌套
	//[ [0, 0, 0], [0, 0, 0] ]
	//var nestData [2][3]int
	//nestData[0] = [3]int{11, 22, 33}
	//nestData[1] = [3]int{44, 55, 66}
	nestData := [2][3]int{[3]int{11, 22, 33}, [3]int{44, 55, 66}}
	fmt.Println(nestData)
	nestData[1][1] = 621
	fmt.Println(nestData)
}
