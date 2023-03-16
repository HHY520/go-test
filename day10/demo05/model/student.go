package model

type student struct {
	Name string
	Score float64
}

// student 结构体是私有的
// 通过工厂模式解决这个问题

func NewStudent(name string,score float64)(*student){
	return &student{
		Name:name,
		Score:score,
	}
}