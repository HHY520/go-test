package main

import "fmt"

func main(){
	var price float32 = 89.23
	fmt.Println(price)

	var price1 float32 = -123.0000901
	var price2 float64 = -123.0000901
	fmt.Println(price1, price2)

	var price3 = 1.12
	fmt.Printf("price3的数据类型为%T \n", price3)

	price4 := 0.123
	fmt.Printf("price4的数据类型为%T", price4)
}