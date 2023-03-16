package main

import (
	"fmt"
	"math/rand"
	"time"
)

func test(arr *[3]int){
	fmt.Printf("1111 %p \n",arr)
	fmt.Printf("3333 %p \n",&arr)
	arr[1] = 20
	fmt.Println("1111",*arr)
}

func test1(){
	var arr [26]byte
	// 65 - 90 A - Z ; 97 - 122 a - z
	for i := 0; i < 26 ; i++{
		arr[i] = 'A' + byte(i)
	}
	for _ , val := range arr{
		fmt.Printf("%c ",val)
	}
	fmt.Println(arr)
}

func test2(){
	var arr = [...]int{1,-1,20,30,25,90,100}
	var max int = 0
	var index1 int
	for index ,val :=range arr{
		if max < val {
			max = val
			index1 = index
		}
	}
	fmt.Println(max,index1)
}

func test3(){
	var arr = [...]int {1,2,4,5,7,8,9}
	var sum int
	var pinjun float64
	for _ ,val := range arr{
		sum += val
	}
	pinjun = float64(sum) / float64(len(arr))
	fmt.Println(sum,pinjun)
}

func test4(){
	var arr [5] int
	for i := 0 ; i < len(arr) ; i++{
		rand.Seed(time.Now(). UnixNano())
		arr [i] = rand.Intn(100)
	}
	fmt.Println(arr)
	var arr1 [5] int
	for index , val := range arr{
		arr1[len(arr) - index - 1] = val
	}
	fmt.Println(arr1)
	
}

func main() {
	test4()
}

