package main

import (
	"fmt"
)

// 斐波那契数列练习
func test1(n int){
	var slice []uint = make([]uint,n)
	for i := 1 ; i <= n ; i++{
		if i == 1 || i == 2{
			slice[i-1] = 1
		}else{
			slice[i-1] = slice[i-3] + slice[i-2]
		}
	}
	fmt.Println(slice)
}

func main() {
	test1(10)
}