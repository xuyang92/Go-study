package main

import "fmt"

func main() {
	//for循环嵌套 + continue
f1:
	for i := 1; i < 3; i++ {
		for j := 1; j < 5; j++ {
			if j == 3 {
				continue f1
			}
			fmt.Println(i, j)
		}
	}

f2:
	for i := 1; i < 3; i++ {
		for j := 1; j < 5; j++ {
			if j == 3 {
				break f2
			}
			fmt.Println(i, j)
		}
	}
}
