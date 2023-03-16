package main

import (
	"fmt"
)
type A struct {
	Name string
	Age int
}

type B struct {
	Name string
	Score int
}

type C struct {
	Name string
	A
	*B
	int
}

func main(){
	var c = C{A:A{Name:"Tom",Age:18},B:&B{Name:"Tina",Score:90},Name:"haha"}
	fmt.Println(c) // {haha {Tom 18} 0x1400011c018}
	fmt.Println(*c.B) //{Tina 90}
}





