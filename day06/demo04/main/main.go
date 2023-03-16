package main

import (
	"fmt"
	"day06/demo04/pkg"
)

func main() {
	var a int = 3
	var b int = 2
	var operator byte = '*'
	fmt.Println(pkg.Cal(a,b,operator))
	
}