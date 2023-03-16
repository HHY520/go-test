package main

import(
	"fmt"
)

func main() {
	var a int = 1
	var b int = 2
	a,b = b,a
	fmt.Println(a,b)

}