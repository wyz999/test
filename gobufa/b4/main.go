package main

import (
	"fmt"
	"sync"
)

//并发版本的hello word
//使用
func wg() {

	var wg sync.WaitGroup
	//开启n个打印线程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Printf("hello word+ %d \n", i)
			wg.Done()
		}(i)
	}
	//等待n个线程完成
	wg.Wait()
}

func ch() {
	ch := make(chan int, 10)
	for i := 0; i < cap(ch); i++ {
		go func(i int) {
			ch <- i
			fmt.Printf("hello word -%d \n", i+100-1)
		}(i)
	}
	for i := 0; i < cap(ch); i++ {
		<-ch
	}
}

func main() {
	wg()
	ch()
}
