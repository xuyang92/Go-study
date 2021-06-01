package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// map
func main() {
	/*
	var m1 map[string]int
	fmt.Println(m1 == nil) 	//还没有初始化（没有在内存中开辟空间）
	//要估算好该map容量，避免在程序运行期间再动态扩容
	m1 = make(map[string]int, 10)
	m1["理想"] = 18
	m1["梦想"] = 99
	m1["白日梦"] = 250
	fmt.Println(m1)
	v, ok := m1["程序员"]
	if !ok {
		fmt.Println("没有程序员")
	} else {
		fmt.Println(v)
	}

	// map的遍历
	for key, value := range m1 {
		fmt.Println(key, value)
	}
	for key := range m1 {
		fmt.Println(key)
	}
	for _, value := range m1 {
		fmt.Println(value)
	}

	// 删除
	delete(m1, "梦想")
	fmt.Println(m1)
	*/

	//****** 按照指定顺序遍历map
	// 获取当前时间的纳秒数
	rand.Seed(time.Now().UnixNano())	//初始化随机数种子.
	scoreMap := make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)			 // 生成0~99的随机整数
		scoreMap[key] = value
	}
	// 取出map中的所有key存入切片keys
	keys := make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	// 对切片进行排序
	sort.Strings(keys)
	// 按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
