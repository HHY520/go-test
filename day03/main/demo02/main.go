package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// var i int = 1
	// fmt.Println(i)
	// var n2 = 100
	// fmt.Printf("n2 的 类型 %T" , n2)

	// var num = 123
	// var num1 int = 234
	// var num2 string = "456"
	// num3 := 345
	// fmt.Println(num, num1, num2, num3)

	// var (
	// 	num int
	// 	num1 = 1
	// 	num2 int = 2
	// )
	// fmt.Println(num, num1, num2)
	var a, b = GetData(10,20)
	fmt.Println(a, b)

	var c = 100
	fmt.Printf("c 的类型是 %T", c)
	fmt.Println()

	fmt.Printf("c 的类型是 %T  c占用的字节数%d", c, unsafe.Sizeof(c))

}

func GetData(num int,num1 int) (int, int){
	return num,num1
}