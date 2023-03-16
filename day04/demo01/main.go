package main

import(
	"fmt"
)

func main() {
	// 算术运算符
	// fmt.Println(float32(10)/float32(4))
	// fmt.Println(10%3)
	// var a int = 1
	// a++
	// var b = a
	// fmt.Println(b)
	// fmt.Println(97/7,"星期")
	// fmt.Println("零",97%7,"天")

	// 关系运算符
	// fmt.Println(10==10)
	// fmt.Println(10!=10)
	// fmt.Println(1>2)
	// fmt.Println(1<2)

	//逻辑运算符
	if true && false{
		fmt.Println(1)
	}else{
		fmt.Println(2)
	}

	if true || false{
		fmt.Println(1)
	}else{
		fmt.Println(2)
	}

	if !(true && false){
		fmt.Println(1)
	}else{
		fmt.Println(2)
	}





}