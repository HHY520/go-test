package main

import (
	"fmt"
)

type Student struct {
	Name string
	Gender string
	Age int
	Id int
	Score float64
}

func (stu *Student) say()(string){
	return fmt.Sprintf("name是:%v；gender是:%v；age是:%v；id是%v;score是%v",stu.Name,stu.Gender,stu.Age,stu.Id,
	stu.Score)
}

type Visitor struct {
	Name string
	Age int
}

func (vis *Visitor) getTickets() {
	if vis.Age > 18 {
		fmt.Println("姓名：",vis.Name)
		fmt.Println("年龄：",vis.Age)
		fmt.Println("收费：20元")
	}else {
		fmt.Println("姓名：",vis.Name)
		fmt.Println("年龄：",vis.Age)
		fmt.Println("收费：免费")
	}
}

func main(){
	var student Student = Student{"Tom","男",18,1,99.9}
	fmt.Println(student.say())

	var visitor Visitor = Visitor{"Tina",19}
	visitor.getTickets()
}

