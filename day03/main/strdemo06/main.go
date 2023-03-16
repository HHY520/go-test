package main

import "fmt"

func main() {
	var str1 string = "上海"
	fmt.Println(str1)

	var str2 string = "上海" +
	"你好" +
	"哈哈"
	fmt.Println(str2)

	var a int
	var b float32
	var c float64
	var d bool
	var e string
	fmt.Println("int 默认值",a,"float32 默认值",b,"float64 默认值",c,"bool 默认值",d,"string 默认值",e)
}