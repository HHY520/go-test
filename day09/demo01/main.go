package main

import (
	"fmt"
)

func test1() {
	// map 的声明时,直接赋值
	var a = make(map[string]map[string]string)
	a["1"] = make(map[string]string,2)
	a["1"]["name"] = "tom"
	a["1"]["sex"] = "男"
	a["2"] = make(map[string]string,2)
	a["2"]["name"] = "tina"
	a["2"]["sex"] = "女"
	fmt.Println(a)

	// val, findRes := a["2"]
	// if findRes{
	// 	fmt.Println("找到了",val)
	// }else{
	// 	fmt.Println("没有找到")
	// }

	for key , value := range a{
		fmt.Println("key=",key)
		fmt.Println("value",value)
	}
	fmt.Println(len(a))
}

func main() {
	// 定义一个 map 切片
	var sliceMap = make([]map[string]string, 2)
	// 增加 map 信息
	if sliceMap[0] == nil{
		sliceMap[0] = make(map[string]string, 2)
		sliceMap[0]["name"] = "牛魔王"
		sliceMap[0]["age"] = "500"
	}
	// 增加 map 信息
	if sliceMap[1] == nil{
		sliceMap[1] = make(map[string]string, 2)
		sliceMap[1]["name"] = "红孩儿"
		sliceMap[1]["age"] = "400"
	}
	// 增加 map 信息（错误：数组越界）
	// if sliceMap[2] == nil{
	// 	sliceMap[2] = make(map[string]string, 2)
	// 	sliceMap[2]["name"] = "狐狸精"
	// 	sliceMap[2]["age"] = "200"
	// }
	
	// 使用 append 函数动态增加切片
	newMonster := map[string]string{
		"name" : "狐狸精",
		"age" : "200",
	}
	sliceMap = append(sliceMap,newMonster)
	slice := make(map[string]string, 2)
	slice["name"] = "红孩儿"
	slice["age"] = "200"
	sliceMap = append(sliceMap,slice)
	fmt.Println(sliceMap)
	
}

