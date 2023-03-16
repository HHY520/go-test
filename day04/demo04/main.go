package main

import(
	"fmt"
)

func main() {
	var name string
	var age byte
	var sal float32
	var isPass bool
	// 一个一个输入
	// fmt.Println("请输入姓名:")
	// fmt.Scanln(&name)

	// fmt.Println("请输入年龄:")
	// fmt.Scanln(&age)

	// fmt.Println("请输入薪水:")
	// fmt.Scanln(&sal)

	// fmt.Println("请输入是否通过考试:")
	// fmt.Scanln(&isPass)

	// fmt.Println(name,age,sal,isPass)

	// 一次性输入
	fmt.Println("请输入姓名、年龄、薪水、是否通过考试，使用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Println(name,age,sal,isPass)


}