package main

import(
	"fmt"
)

func main() {
	// 短路与
	if false && true{
		fmt.Println("1")
	}else{
		fmt.Println("2")
	}
	if true && true{
		fmt.Println("1")
	}else{
		fmt.Println("2")
	}

	// 短路或
	if true || false{
		fmt.Println("1")
	}else{
		fmt.Println("2")
	}
	if false || false{
		fmt.Println("1")
	}else{
		fmt.Println("2")
	}

}