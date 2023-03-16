package main

import (
	"fmt"
	"sort"
)

func main(){
	// map 的排序
	map1 := make(map[int]int,10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 10
	map1[8] = 20
	fmt.Println(map1)
	// 如何排序 kay
	// 1.先讲 map key 放入到切片中
	// 2.对切片进行排序
	// 3.遍历切片，然后按照 key 来输出排序

	var keys []int
	for key,_ := range map1{
		keys = append(keys,key)
	}
	// 排序
	sort.Ints(keys)
	for _,key := range keys{
		fmt.Println(key," = ",map1[key])
	}
}