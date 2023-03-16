package main 

import (
	"fmt"
)

func main () {
	// 练习 1
	// var sum int;
	// for i := 1 ; i <= 100 ; i++ {
	// 	if i % 9 == 0 {
	// 		sum += i
	// 	}
	// }
	// fmt.Println("sum = ",sum)

	// 练习 2
	// for i := 6 ; i >= 0 ; i-- {
	// 	fmt.Println(6 - i ," + ",i," = 6")

	// }

	// 练习1：几个班学生成绩平均分

	// var class int = 2
	// var student int = 5 
	// var sumClass float64 = 0
	
	// for i := 1 ; i <= class ; i++{
	// 	fmt.Printf("请开始输入第 %v 个班的学生成绩：\n" , i)
	// 	for j:= 1; j<= student ; j++{
	// 		fmt.Printf("请输入第 %v 个班级，第 %v 个学生的成绩：\n" , i , j)
	// 		var kay float64;
	// 		fmt.Scanln(&kay)
	// 		sumClass += kay
	// 	}
	// 	fmt.Printf("总成绩：%v，平均分是：%v \n", sumClass , sumClass / float64(i * student))
	// }

	// 练习2：统计几个班学生及格人数

	// var class int = 2
	// var student int = 5 
	// var sumStudent float64 = 0
	
	// for i := 1 ; i <= class ; i++{
	// 	fmt.Printf("请开始输入第 %v 个班的学生成绩：\n" , i)
	// 	for j:= 1; j<= student ; j++{
	// 		fmt.Printf("请输入第 %v 个班级，第 %v 个学生的成绩：\n" , i , j)
	// 		var kay float64;
	// 		fmt.Scanln(&kay)
	// 		if kay >= 60{
	// 			sumStudent += 1
	// 		}
	// 	}
	// 	fmt.Printf("及格人数：%v \n", sumStudent)
	// }

	// 练习3：打印金字塔【实心、空心、空心菱形】
	
	// 设置打印金字塔的行数
	// var total int = 20

	// for i := 1 ; i <= total ; i++{
	// 	// 打印 ” “ 程序
	// 	for k := 1 ; k <= total - i ; k++{
	// 		fmt.Printf(" ")
	// 	}
	// 	// 打印 * 程序
	// 	for j := 1 ; j <= 2 * i - 1 ; j++{
	// 		if (j > 1 && j < 2 * i - 1) && i != total{
	// 			fmt.Printf(" ")
	// 		}else{
	// 			fmt.Printf("*")
	// 		}
	// 	}
	// 	fmt.Println()
	// }


	// 练习4：打印九九乘法表
	// var total int = 9

	// for i := 1 ; i <= total ; i++{
	// 	for j := 1 ; j <= i ; j++{
	// 		fmt.Printf("%v * %v = %v \t" , j , i , j*i)
	// 	}
	// 	fmt.Println()
	// }


}