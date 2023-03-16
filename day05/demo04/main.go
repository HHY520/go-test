package main

import (
	"fmt"
	_"math/rand"
	_"time"
)

func main() {
	// break 正常使用
	// var count int = 0
	// for{
	// 	rand.Seed(time.Now().UnixNano())
	// 	var randNum int = rand.Intn(100) + 1
	// 	count++
	// 	if randNum == 99{
	// 		break
	// 	}
	// }
	// fmt.Println("生成 99 一共使用了：",count)

	// break 指定标签跳出
	// lable1:
	// for i := 0 ; i <= 4 ; i++{
	// 	// lable2:
	// 	for j := 0 ; j <= 10 ; j ++{
	// 		if j == 2 {
	// 			// break // break 默认会跳出最近的 for 循环
	// 			// break lable2 // break lable2 会跳出 lable2 循环
	// 			break lable1 // break lable1 会跳出 lable1 循环
	// 		}
	// 		fmt.Println("j=",j)
	// 	}
	// }
	
	// 练习1:100以内求和，和第一次大于20的当前数

	// 定义一个求和变量
	// var sumNumber int = 0
	// i := 0
	// for ; i <= 100 ; i++{
	// 	sumNumber += i
	// 	if sumNumber > 20{
	// 		break
	// 	}else{
	// 		continue
	// 	}
	// }
	// fmt.Println(i)

	// 实现登录验证，有三次机会，成功提示登录成功，失败提示还有几次机会

	// 定义用户账号
	var username string = "张三"
	var password string = "123456"
	// 定义输入次数
	var count int = 3
	for i := 1 ; i <= count ; i++{
		fmt.Printf("第 %v 次输入账号密码，还剩余 %v 次\n",i,count-i)
		var name string
		var pass string
		fmt.Println("请输入用户账号：")
		fmt.Scanln(&name)
		fmt.Println("请输入用户密码：")
		fmt.Scanln(&pass)
		if username == name && password == pass{
			fmt.Println("登录成功")
			break
		}else{
			fmt.Println("密码输入错误，请重新尝试！！！")
			continue
		}
	}
	
}