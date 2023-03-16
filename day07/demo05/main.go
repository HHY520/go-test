package main

import (
	"fmt"
	"strings"
)

// 累加器
func AddUpper() func (int) int {
	var n int = 10
	return func (x int) int {
		n = n + x
		return n
	}
}

func makeSuffix(suffix string) func (string) string {
	return func (fileName string) string {
		if strings.HasSuffix(fileName,suffix) {
			return fileName
		} else{
			return fileName + suffix
		}
	}
}

func main() {
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))

	f1 := makeSuffix(".jpg")
	fmt.Println(f1("haha"))
}
