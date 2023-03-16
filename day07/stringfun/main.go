package main

import (
	"fmt"
	"strconv"
	_"strings"
	"time"
)

func test03(){
	str := ""
	for i := 0 ; i < 100000 ; i ++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func main() {
	now1 := time.Now().UnixNano()
	test03()
	now2 := time.Now().UnixNano()
	fmt.Println(now2-now1)
}
