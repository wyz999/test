package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//原子操作
var total uint64

func worker(gw *sync.WaitGroup) {
	defer gw.Done()
	var i uint64
	for i = 0; i <= 100; i++ {
		//atomic.AddUint64函数调用保证了total的读取、更新和保存是一个原子操作，因此在多线程中访问也是安全的
		atomic.AddUint64(&total, i)
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()
	fmt.Println(total) //10100
}
