package main

import (
	"fmt"
)

func modifyUser(users map[string]map[string]string,name string){
	for _,value := range users{
		if value["nickname"] == name{
			value["pwd"] = "888888"
		}
	}
}

func main(){
	users := make(map[string]map[string]string)
	user1 := make(map[string]string)
	user1["nickname"] = "tina"
	user1["pwd"] = "123456"
	users["no1"] = user1
	user2 := make(map[string]string)
	user2["nickname"] = "mary"
	user2["pwd"] = "6789"
	users["no2"] = user2
	fmt.Println(users)
	modifyUser(users,"tina")
	fmt.Println(users)

}