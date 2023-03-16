package utils

import (
	"fmt"
)

func Cal(a int,b int,operator byte)(int){
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
		fmt.Println("操作符有误！！！")
	}
	return res
}