package model

import (
	"fmt"
)

type accounts struct {
	account string
	balance float64
	password string
}

func NewAccount(account string) *accounts{
	if len(account) >= 6 && len(account) <= 10{
		return &accounts{
			account : account,
		}
	} else {
		fmt.Println("账户输入长度不合规")
		return nil
	}
}

// 对密码、余额进行封装处理
func (a *accounts) SetBalance(balance float64){
	if balance >= 20 {
		a.balance = balance
	} else {
		fmt.Println("余额不合规")
	}
}

func (a *accounts) GetBalance()float64{
	return a.balance
}


func (a *accounts) SetPassword(password string){
	if len(password) ==6 {
		a.password = password
	} else {
		fmt.Println("密码输入长度不合规")
	}
}

func (a *accounts) GetPassword()string{
	return a.password
}