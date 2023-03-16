package main

import (
	"fmt"
)
var (
	Fun1 = func(n1 int ,n2 int) int{
		return n1 + n2
	}
)

func main() {
	a := Fun1(10,20)
	fmt.Println(a)
}