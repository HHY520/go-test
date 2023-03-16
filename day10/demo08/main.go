package main

import (
	"fmt"
)
type A struct {
	Name string
	age int
}
func (a *A) SayOK() {
	fmt.Println("A SayOk",a.Name)
}
func (a *A) hello(){
	fmt.Println("A hello",a.Name)
}

type B struct {
	A
}

func main() {
	var b B
	b.A.Name = "Tom"
	b.A.age = 18
	b.A.SayOK()
	b.A.hello()

	b.Name = "Marry"
	b.age = 20
	b.SayOK()
	b.hello()
}

