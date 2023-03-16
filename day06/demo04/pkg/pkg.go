package pkg

import (
	"fmt"
)

func Cal (a int,b int,operator byte)(int) {
	var res int
	switch operator{
		case '+':
			res = a + b
		case '-':
			res = a - b
		case '*':
			res = a * b
		case '/':
			res = a / b
		default:
			fmt.Println("输入计算符号有问题！！！")
	}
	return res
}