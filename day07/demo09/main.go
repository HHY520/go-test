package main

import (
	"fmt"
	"errors"
)

func readConf(name string)(err error){
	if name == "config.ini"{
		return nil
	}else{
		// 返回一个自定义错误
		return errors.New("读取文件错误")
	}
}

func test02(){
	err := readConf("config.in")
	if err != nil {
		//如果文件读取失败，就输出错误，并终止程序
		panic(err)
	}
	fmt.Println("11111")
}

func main() {
	test02()
}
