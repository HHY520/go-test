package main

import (
	"fmt"
)

func test(n int)(int) {
		num := 0
		if n == 1 || n == 2{
			num = 1
		} else{
			num = test(n-1) + test(n-2)
		}
		return num
}

func test2(n int)(int) {
	if n == 1{
		return 3
	}else{
		return 2 * test2(n-1) + 1
	}
}

func test3(n int)(int){
	if n == 10{
		return 1
	}else{
		return (test3(n + 1) + 1) * 2
	}
}

func test4(n1 *int,n2 *int){
	*n1,*n2 = *n2,*n1
}


func main() {
	fmt.Println(test(6))
	fmt.Println(test2(5))
	fmt.Println(test3(1))
	n1 := 3
	n2 := 4
	test4(&n1,&n2)
	fmt.Println("n1=",n1,"n2=",n2)
}