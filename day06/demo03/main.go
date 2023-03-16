package main

import (
	"fmt"
)

func test(n1 int){
	n1 += 1
	fmt.Println(n1)
}

func getSum(n1 int,n2 int)(int){
	return n1 + n2
}

func getSumAndSub (n1 int,n2 int)(int , int){
	return n1+n2 , n1-n2
}

func main(){
	n1 := 0

	test(n1)
	fmt.Println("n1=",n1)
	fmt.Println(getSum(1,2))

	rest1 , rest2 := getSumAndSub(2,4)
	fmt.Println(rest1,rest2)
}