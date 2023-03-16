package main

import (
	"fmt"
)

func main() {
	// 练习一
	// var charType byte
	// fmt.Println("请输入一个小写字母：")
	// fmt.Scanf("%c",&charType)
	// switch charType {
	// 	case 'a':
	// 		fmt.Println("A")
	// 	case 'b':
	// 		fmt.Println("B")
	// 	case 'c':
	// 		fmt.Println("C")
	// 	case 'd':
	// 		fmt.Println("D")
	// 	case 'e':
	// 		fmt.Println("E")
	// 	default:
	// 		fmt.Println("other")
	// }

	// 练习二
	var kay int
	fmt.Println("请输入学生的成绩")
	fmt.Scanln(&kay)
	switch{
		case kay > 60 && kay <= 100:
			fmt.Println("合格")
		case kay <= 60 && kay >= 0:
			fmt.Println("不合格")
		default:
			fmt.Println("输入有误")
	}

}
