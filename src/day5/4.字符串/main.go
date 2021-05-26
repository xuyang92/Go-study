package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	// 1、本质是utf-8编码的序列
	var name string = "曹煦阳"
	fmt.Println(name[0], strconv.FormatInt(int64(name[0]), 2))
	fmt.Println(name[1], strconv.FormatInt(int64(name[1]), 2))
	fmt.Println(name[2], strconv.FormatInt(int64(name[2]), 2))

	fmt.Println(name[3], strconv.FormatInt(int64(name[3]), 2))
	fmt.Println(name[4], strconv.FormatInt(int64(name[4]), 2))
	fmt.Println(name[5], strconv.FormatInt(int64(name[5]), 2))

	fmt.Println(name[6], strconv.FormatInt(int64(name[6]), 2))
	fmt.Println(name[7], strconv.FormatInt(int64(name[7]), 2))
	fmt.Println(name[8], strconv.FormatInt(int64(name[8]), 2))

	// 2、获取字符串的长度：9（字节长度）
	fmt.Println(len(name))

	// 3、字符串转换为一个“字节集合”
	byteSet := []byte(name)
	fmt.Println(byteSet)

	// 4、字节的集合转换为字符串
	byteList := []byte{230, 155, 185, 231, 133, 166, 233, 152, 179}
	targetString := string(byteList)
	fmt.Println(targetString)

	// 5、rune的集合:将字符串转换为unicode字符集码点的集合
	tempSet := []rune(name)
	fmt.Println(tempSet)
	fmt.Println(tempSet[0], strconv.FormatInt(int64(tempSet[0]), 16))
	fmt.Println(tempSet[1], strconv.FormatInt(int64(tempSet[1]), 16))
	fmt.Println(tempSet[2], strconv.FormatInt(int64(tempSet[2]), 16))

	// 6、rune集合转换为字符串
	runeList := []rune{26361, 29030, 38451}
	targetName := string(runeList)
	fmt.Println(targetName)

	// 7、长度的处理
	runeLength := utf8.RuneCountInString(name)
	fmt.Println(runeLength)
}
