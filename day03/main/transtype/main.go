package main

import "fmt"
func main() {
	var i int = 100
	var j float32 = float32(i)
	fmt.Println(j)
	var c string = fmt.Sprintf("%d",i)
	fmt.Println(c)

}