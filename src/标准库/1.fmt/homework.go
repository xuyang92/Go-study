package main

import (
	"fmt"
)

// 你有50枚金币，需要分配给以下几个人：Matthew，Sarah，Augustus，Heidi，
// Emilie,Peter,Giana,Adriano,Aaron,Elizabeth
// 分配规则如下：
// a.名字中每包含1个'e'或'E'分1枚金币
// b.名字中每包含1个'i'或'I'分2枚金币
// c.名字中每包含1个'o'或'O'分3枚金币
// d.名字中每包含1个'u'或'U'分4枚金币
// 写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
// 程序结构如下：请实现'dispatchCoin‘函数

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}

	distribution = make(map[string]int, len(users))
)

func dispatchCoin() int {
	// 依次拿到每个人的名字;
	for _, name := range users {
		// 拿到一个人名字根据分金币的规则法分金币
		for _, c := range name {
			switch c{
			case 'e', 'E':
				coins -= 1
				distribution[name] ++
			case 'i', 'I':
				coins -= 2
				distribution[name] += 2
			case 'o', 'O':
				coins -= 3
				distribution[name] += 3
			case 'u', 'U':
				coins -= 4
				distribution[name] += 4
			}
		}
	}
	return coins
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
	fmt.Println(distribution)
}
