package main

import (
	"fmt"
	"day10/demo05/model"
)


func main() {
	var stu = model.NewStudent("Tom",78.9)
	fmt.Println((*stu).Name,(*stu).Score)
}

