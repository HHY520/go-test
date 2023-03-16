package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 练习 1 
func test01 (){
	for {
		var year string
		fmt.Println("请输入年份：")
		fmt.Scanln(&year)
		var mounth int
		fmt.Println("请输入月份：")
		fmt.Scanln(&mounth)
		if mounth < 0 || mounth > 12 {
			fmt.Println("月份输入有问题！！！")
			continue
		}
		var day int
		fmt.Println("请输入天：")
		fmt.Scanln(&day)
		fmt.Printf("输入的日期为：%v年%v月%v日 \n",year,mounth,day)
	}
}

// 练习 2
func test02 (){
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(100)+1
	fmt.Println(randNum)
	count := 0
	for {
		count ++
		fmt.Println("输入你所猜的数字：")
		var scanlnNum int
		fmt.Scanln(&scanlnNum)
		if scanlnNum == randNum{
			switch count{
				case 2,3:
					fmt.Println("你真是一个天才")
				case 4,5,6,7,8,9:
					fmt.Println("一般般")
				case 10:
					fmt.Println("可算猜对了")
				default:
					fmt.Println("")
			}
			break
		}else {
			fmt.Println("猜错了！！！")
		}
		if count == 10{
			fmt.Println("游戏结束")
			break
		}
	}
	


}

// 练习 3
func test03 (){
	for i := 2 ; i <= 100 ; i++ {
		var flag bool = true
		for j := 2 ; j < i ; j ++{
			if i % j == 0 {
				flag = false
			}
		}
		if flag {
			fmt.Println(i)
		}
	}
}

// 计算两个时间的差值
func GetDaysBetween2Date(format, date1Str, date2Str string) (int, error) {
	// 将字符串转化为Time格式
  date1, err := time.ParseInLocation(format, date1Str, time.Local)
  if err != nil {
	  return 0, err
  }
	// 将字符串转化为Time格式
  date2, err := time.ParseInLocation(format, date2Str, time.Local)
  if err != nil {
	  return 0, err
  }
  //计算相差天数
  return int(date1.Sub(date2).Hours() / 24), nil
}

// 练习 4
func test04 (){
	date , _ := GetDaysBetween2Date("2006-01-02", "1990-01-01", "2022-06-01")
	dateCha := date % 5
	if dateCha <= 3{
		fmt.Println("打鱼")
	}else{
		fmt.Println("筛网")
	}


}

func main() {
	test04()
}