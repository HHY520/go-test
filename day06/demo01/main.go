package main

import (
	"fmt"
)

func cal(a int,b int,operator byte)(int){
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

func main() {
	var a int= 10
	var b int= 20
	var operator byte = '+'
	fmt.Println(cal(a,b,operator))
}

