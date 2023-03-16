package main

import(
	"fmt"
)

func main() {
	var age byte
	fmt.Println("请输入你的年龄：")
	fmt.Scanln(&age)
	if age < 18 {
		fmt.Println("你未满18岁！！")
	} else if age == 18 {
		fmt.Println("你等于18岁！！！")
	}else{
		fmt.Println("你超过18岁！！！")
	}
}