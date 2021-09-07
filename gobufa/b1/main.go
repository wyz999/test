package main

import (
	"fmt"
	"time"
)

var balance int

func Deposit(amount int) {
	balance += amount
}
func Balance() int {
	return balance
}
func main() {
	go func() {
		Deposit(200)
		fmt.Println("dc=", Balance())
	}()
	Deposit(100)
	fmt.Println("de=", Balance())
	time.Sleep(time.Second * 2)
}
