package main

import "fmt"

func main() {
	// 转义字符的使用
	fmt.Println("tom\thaha")
	fmt.Println("tom\nthaha")
	fmt.Println("tom\\thaha")
	fmt.Println("tom\"thaha\"")
	fmt.Println("tomsdadssd\rthaha")
	fmt.Println("姓名\t年龄\t籍贯\t地址\njoin\t12\t河北\t北京")
}