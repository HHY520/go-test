package main

import (
	"fmt"
)

func main() {
	a := new(string)
	fmt.Printf("a 的 type = %T \n",a)
	fmt.Println("a 的值=",*a)
	*a = "haha"
	fmt.Println("a 的值=",*a)
}