package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// 1.判断字符串中汉字的数量
	s1 := "hello,北京欢迎你"
	var count int
	for _, c := range s1 { // 依次拿到字符串中的字符
		if unicode.Is(unicode.Han, c) { // 判断当前这个字符是不是汉字
			// 把汉字出现的次数累加得到最终结果
			count++
		}
	}
	fmt.Println(count)

	// 2.写一个程序，统计一个字符串中每个单词出现的次数，比如："how do you do"中how=1 do=2 you=1
	// 把字符串按照空格切割得到切片
	s2 := "how do you do"
	s3 := strings.Split(s2, " ")
	// 遍历切片存储到一个map
	m1 := make(map[string]int, 10)
	for _, w := range s3 { // 累加出现的次数
		// 如果原来map中不存在w这个key，那么出现次数等于1
		// 如果map中存在这个key，那么出现次数+1
		if _, ok := m1[w]; !ok {
			m1[w] = 1
		} else {
			m1[w]++
		}
	}
	for key, value := range m1 {
		fmt.Println(key, value)
	}

	// 3.回文判断（字符串从左往右读和从右往左读是一样的，那么就是回文）
	// 例子：上海自来水来自海上、山西运煤车煤运西山、黄山落叶松叶落山黄
	ss := "上海自来水来自海上"
	ret := true
	// 把字符串中的字符拿出来放到一个[]rune中
	r := make([]rune, 0, len(ss))
	for _, c := range ss {
		r = append(r, c)
	}
	for i := 0; i < len(r) / 2; i++ {
		if r[i] != r[len(r)-1-i] {
			ret = false
			break
		}
	}
	if ret {
		fmt.Println("是回文")
	} else {
		fmt.Println("不是回文")
	}
}
