package main

import (
	"fmt"
	ut "day06/demo02/utils"
)

func main() {
	var a int= 10
	var b int= 20
	var operator byte = '+'
	fmt.Println(ut.Cal(a,b,operator))
}