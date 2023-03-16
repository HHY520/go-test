package main

import (
	"fmt"
)

func main() {
	lable2:
	for i := 0 ; i < 4 ; i++ {
		// lable1:
		for j := 0 ; j < 4 ; j++{
			if j == 2{
				// continue // continue 默认结束本次循环，开始下一次循环
				// continue lable1 // continue lable1 结束 lable1 本次循环，继续 lable1 循环
				continue lable2 // continue lable2 结束 lable2 本次循环，继续 lable2 循环
			}
			fmt.Println(j)
		}
	}
}
