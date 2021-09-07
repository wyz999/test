package main

import (
	"fmt"
	"time"
)

func fib(n int) int {
	if n<2 {
		return 1
	}
	return fib(n-1)+fib(n-2)
}

func daySleep(dur time.Duration)  {
	for  {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(dur)
		}
	}
}

func main() {
	n:=45
	go daySleep(2*time.Second)
	fmt.Println(fib(n))
}
