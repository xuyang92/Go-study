package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	//	1、获取字符串长度;
	var name string = "曹煦阳"
	fmt.Println(len(name))						//获取字节长度,输出9
	fmt.Println(utf8.RuneCountInString(name))	//获取字符长度，输出3

	// 2、是否以**开头
	result := strings.HasPrefix(name, "曹")
	fmt.Println(result)		//输出true

	// 3、是否以xx结尾
	result = strings.HasSuffix(name, "阳")
	fmt.Println(result)		//输出true

	// 4、是否包含
	var plan string = "小明计划明天上午九点去游泳"
	result = strings.Contains(plan, "计划")
	fmt.Println(result)		//输出true

	// 5、变大写
	var englishName string = "caoxuyang"
	upEnglishName := strings.ToUpper(englishName)
	fmt.Println(upEnglishName)

	// 6、变小写
	lowerEnglishName := strings.ToLower(upEnglishName)
	fmt.Println(lowerEnglishName)

	// 7、去两边
	result1 := strings.TrimRight(plan, "游泳")	//去除右边
	result2 := strings.TrimLeft(plan, "小")		//去除左边
	result3 := strings.Trim(plan, "去游泳")		//去除两边
	fmt.Println(result1, result2, result3)

	// 8、替换
	var oneDay string = "小明计划明天上午九点去游泳，然后去打球，然后去坐公交回家"
	//找到【去】替换为【完成】，从左到右找第一个替换
	result4 := strings.Replace(oneDay, "去", "完成", 1)
	//找到【去】替换为【完成】，从左到右找前两个替换
	result5 := strings.Replace(oneDay, "去", "完成", 2)
	//找到【去】替换为【完成】，替换所有
	result6 := strings.Replace(oneDay, "去", "完成", -1)
	fmt.Println(result4, result5, result6)

	// 9、分割
	splitList := strings.Split(oneDay, "去")
	fmt.Println(splitList)

	// 10、拼接
	//不建议
	message := "我爱" + "北京天安门"
	fmt.Println(message)

	//建议：效率高一些
	messageList := []string{"我爱", "北京天安门"}
	message1 := strings.Join(messageList, "")
	fmt.Println(message1)

	//建议：效率更高一些（go1.10之前）
	var buffer bytes.Buffer
	buffer.WriteString("我爱")
	buffer.WriteString("北京天安门")
	message3 := buffer.String()
	fmt.Println(message3)

	//建议：效率更更更高一些（go1.10之后）
	var builder strings.Builder
	builder.WriteString("我爱")
	builder.WriteString("北京天安门")
	message4 := builder.String()
	fmt.Println(message4)

	// 11、string转换为int
	num := "666"
	dataNum, _ := strconv.Atoi(num)
	fmt.Println(dataNum)

	// 12、int转换为string
	stringNum := strconv.Itoa(888)
	fmt.Println(stringNum)

	// 13、字符串和字节集合
	// 字符串转换为字节的集合
	byteSet := []byte(name)
	fmt.Println(byteSet)
	// 字节的集合转换为字符串
	byteList := []byte{230, 155, 185, 231, 133, 166, 233, 152, 179}
	targetString := string(byteList)
	fmt.Println(targetString)

	// 14、字符串和rune集合
	// 字符串转换为rune集合
	tempSet := []rune(name)
	fmt.Println(tempSet)
	fmt.Println(tempSet[0], strconv.FormatInt(int64(tempSet[0]), 16))
	fmt.Println(tempSet[1], strconv.FormatInt(int64(tempSet[1]), 16))
	fmt.Println(tempSet[2], strconv.FormatInt(int64(tempSet[2]), 16))
	// rune集合转换为字符串
	runeList := []rune{26361, 29030, 38451}
	targetName := string(runeList)
	fmt.Println(targetName)

	// 15、string和字符
	//数字转字符串
	v1 := string(65)
	fmt.Println(v1)
	v2 := string(26361)
	fmt.Println(v2)
	//字符串转数字
	v3, size := utf8.DecodeRuneInString("A")
	fmt.Println(v3, size)   //65  1   ...size表示最短可表示的字节
	v4, size := utf8.DecodeRuneInString("曹")
	fmt.Println(v4, size)   //26361  3
}
