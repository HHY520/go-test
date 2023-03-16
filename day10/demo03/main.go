package main

import (
	"fmt"
)

type student struct {
	Name string
	Age int
}

func (stu *student) String() string{
	str := fmt.Sprintf("Name=[%v] Age=[%v]" , stu.Name , stu.Age)
	return str
}

func main(){
	stu := student{
		Name:"Tom",
		Age:20,
	}
	fmt.Println(&stu)
}


