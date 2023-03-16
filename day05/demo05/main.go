package main

import (
	"fmt"
)

func main() {
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Float32())
	// fmt.Println(time.Now().Weekday())
	
	fmt.Println("OK1")
	goto label1
	fmt.Println("OK2")
	fmt.Println("OK3")
	label1:
	fmt.Println("OK4")
	fmt.Println("OK5")
	fmt.Println("OK6")
}
