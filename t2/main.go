package main

import (
	"fmt"
)

func handErr(err error,msg string)  {
	if err != nil {
		fmt.Println(msg+" ",err)
	}
}

type food interface {
	getName() string
}

type apple struct {
	name string
}

type potato struct {
	name string
}

func (a *apple) getName() string {
	a.name="烂苹果"
	return a.name
}

func (p *potato) getName() string {
	p.name="刮的白起"
	return p.name
}

func testRecover() int {
	defer func() {
		fmt.Println(recover())
	}()
	b:=1
	c:=0
	a:=b/c
	fmt.Println(a)
	return 0
}

func main() {
	//使用Go语言中的unsafe.Alignof可以返回相应类型的对齐系数
	//fmt.Printf("int %d\n",unsafe.Alignof(int(0)))
	//fmt.Printf("int8 %d\n",unsafe.Alignof(int8(0)))
	//fmt.Printf("int16 %d\n",unsafe.Alignof(int16(0)))
	//fmt.Printf("int32 %d\n",unsafe.Alignof(int32(0)))
	//fmt.Printf("complex128 %d\n",unsafe.Alignof(complex128(0)))

	var f food
	f=new(apple)
	c:=f.getName()
	f=new(potato)
	w:=f.getName()
	fmt.Printf("s %s \n",c)
	fmt.Printf("w %s",w)

	testRecover()

	//fmt.Println(1<<2)
	//fmt.Println("移动",4>>1)
	//fmt.Println("移动2",16>>2)
}

