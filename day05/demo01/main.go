package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("你好！！！")
	}

	j := 1
	for j < 10 {
		fmt.Println("你好！！！！！")
		j++
	}

	// 死循环，需要与 break 结合使用
	k := 1
	for {
		if k <=10 {
			fmt.Println("哈哈")
		} else {
			break
		}
		k ++
	}
}

