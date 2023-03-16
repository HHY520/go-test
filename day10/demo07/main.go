package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age int
	Score int
}
// 构造共有方法
// 显示学生成绩
func (stu *Student) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v \n",stu.Name,stu.Age,stu.Score)
}

// 设置学生成绩
func (stu *Student) SetScore(score int) {
	stu.Score = score
}

// 编写一个小学生考试系统
type Pupil struct {
	Student
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试中......")
}

// 编写一个大学生考试系统
type Graduate struct {
	Student
}

func (g *Graduate) testing() {
	fmt.Println("大学生正在考试中......")
}



func main() {
	var pupil = &Pupil{}
	pupil.Student.Name = "Tom"
	pupil.Student.Age = 8
	pupil.testing()
	pupil.Student.SetScore(80)
	pupil.Student.ShowInfo()

	var graduate = &Graduate{}
	graduate.Student.Name = "Marry"
	graduate.Student.Age = 18
	graduate.testing()
	graduate.Student.SetScore(90)
	graduate.Student.ShowInfo()
}