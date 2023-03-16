package main

import (
	"fmt"
	"encoding/json"
)

type Monster struct{
	Name string `json:"name"`
	Age int	`json:"age"`
	Skill string `json:"skill"`
}

func main (){
	monster := Monster{"牛魔王",500,"芭蕉扇"}
	
	jsonMonster,err := json.Marshal(monster)
	if err != nil{
		fmt.Println("json 处理错误",err)
	}
	fmt.Println(string(jsonMonster))

}

