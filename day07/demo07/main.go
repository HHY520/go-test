package main

import (
	"fmt"
)

// 当执行 defer 时，暂时不执行，会将 defer 后面的语句压入到独立的栈中（defer 栈）
// 当函数执行完毕后，再从 defer 栈，按照先入后出的方式出栈
func test(n1 *int, n2 *int) int {
	defer fmt.Println("defer关键字 :",*n1)
	defer fmt.Println("1111 :",*n2)
	*n1++
	res := *n1 + *n2
	fmt.Println("OK:",res)
	return res

}

func main() {
	defer fmt.Println("函数执行")
	a := 10
	b := 20
	fmt.Println(test(&a,&b))
}