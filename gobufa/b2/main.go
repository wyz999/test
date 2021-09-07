package main

import (
	"fmt"
	"sync"
	"time"
)

// sync.Mutex互斥锁
//sync.RWMutex读写锁
var (
	Sum int
	s   sync.Mutex
	sm  sync.RWMutex
)

func setSum(n int) {

	s.Lock()
	defer s.Unlock()
	for i := 1; i <= n; i++ {
		Sum += i
	}

}

func getSum() int {
	sm.RLock()
	defer sm.RUnlock()
	return Sum
}

func main() {
	go setSum(6)
	fmt.Println("m1", getSum())
	go setSum(5)
	fmt.Println("m2", getSum())
	time.Sleep(3 * time.Second)
}
