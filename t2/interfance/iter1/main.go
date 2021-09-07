package main

import "fmt"

//接口嵌套
type role interface {
	returnRole() string
}

type people interface {
	role
	getName() string
}


type teacher struct {
	name string
	role string
}

func (t teacher) getName() string {
	return t.name
}

func (t teacher) returnRole() string {
	return t.role
}


func main() {
	var p people
	t:=teacher{
		name: "熊姐",
		role: "辅导员",
	}
	p=t
	fmt.Println("role",p.returnRole())
	fmt.Println("name",p.getName())
}
