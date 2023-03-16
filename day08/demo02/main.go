package main

import (
	"fmt"	
)

func test1(){
	var arr [5]int = [5]int{1,2,3,4,5}
	var slice = arr[1:3]
	fmt.Println("arr=",arr)
	fmt.Println("slice=",slice)
	fmt.Println("slice len=",len(slice))
	fmt.Println("slice cap=",cap(slice))
	// 错误
	// slice[3] = 20
	// fmt.Println("arr=",arr)
	// fmt.Println("slice=",slice)
}

func test2(){
	var slice []int = make([]int,3,5)
	slice1 := make([]int,3,5)
	fmt.Println(slice)
	fmt.Println(slice1)
}

func test3(){
	var slice []int = []int{1,3,5}
	fmt.Println(slice)
}

func test4(){
	var arr []int = []int{1,2,3,4,5}
	var slice = make([]int,1)
	fmt.Println("slice=",slice) // slice= [0]

	copy(slice,arr)
	fmt.Println("slice=",slice) // slice= [1]
}

func main(){
	
}

