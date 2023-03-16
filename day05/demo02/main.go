package main

import (
	"fmt"
)

func main () {
	// 传统方式，遍历字符乱码
	var str string = "hudshfjsdhf"
	for i := 1 ; i < len(str) ; i ++ {
		fmt.Printf("%c \n",str[i])
	}

	var str1 string = "hudshfjsdhf北京"
	// 正常方式，遍历字符正常
	for index , val := range str1 {
		fmt.Printf("index = %d , val = %c \n" , index , val)
	}

	// 遍历汉字，将汉字转化为 []rune 切片
	var str2 string = "hudshfjsdhf哈哈哈"
	str3 := []rune(str2)
	for index , val := range str3 {
		fmt.Printf("index = %d , val = %c \n" , index , val)
	}
	
}
