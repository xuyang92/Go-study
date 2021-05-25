package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	//var v1 float32
	//v1 = 3.14
	//v2 := 99.9	//这样写默认为float64
	//v3 := float64(v1) + v2
	//fmt.Println(v1, v2, v3)
	//
	//v4 := 0.1
	//v5 := 0.2
	//result := v4 + v5
	//fmt.Println(result)
	//
	//v6 := 0.3
	//v7 := 0.2
	//data := v6 + v7
	//fmt.Println(data)
	var v1 = decimal.NewFromFloat(0.000000000000019)
	var v2 = decimal.NewFromFloat(0.29)
	var v3 = v1.Add(v2)
	var v4 = v1.Sub(v2)
	var v5 = v1.Mul(v2)
	var v6 = v1.Div(v2)
	fmt.Println(v3, v4, v5, v6)

	var price = decimal.NewFromFloat(3.4626)
	var data1 = price.Round(1)			//保留小数点后1位（四舍五入）
	var data2 = price.Truncate(1)		//保留小数点后一位（多余位直接舍掉）
	fmt.Println(data1, data2)			//输出3.5  3.4
}
