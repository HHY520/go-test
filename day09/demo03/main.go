package main

import (
	"fmt"
)
// 定义一个结构体
type Stu struct{
	Name string
	Age int
	Address string
}

func main(){
	map1 := make(map[string]Stu,10)
	// 创建学生
	stu1 :=Stu{"tina",12,"北京"}
	stu2 :=Stu{"mary",28,"上海"}
	map1["no1"] = stu1
	map1["no2"] = stu2
	fmt.Println(map1)
	for key,value := range map1{
		fmt.Println(key,"的信息："," name：",value.Name," age：",value.Age," address：",value.Address)
	}
}

