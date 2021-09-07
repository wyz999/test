package main

import (
	"fmt"
	"time"
)

//串联chan
func test1(ch1, ch2 chan int) {
	go func() {
		for i := 0; ; i++ {
			ch1 <- i
		}
	}()

	go func() {
		for {
			x := <-ch1
			ch2 <- x * x
		}
	}()

}

func test2(ch1, ch2 chan int) {
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i * 10

		}
	}()
}

//有缓存的chan
func bufTest(ch1 chan int) {
	go func() {
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
		time.Sleep(time.Second * 5)
		ch1 <- 4
	}()

	go func() {
		time.Sleep(2 * time.Second)
		for i2 := range ch1 {
			fmt.Println("ch ", i2)
		}
	}()
}

//单方向的Channel
func counter(out chan int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}
func squarer(out, in chan int) {
	for i := range out {
		in <- i * i
	}
	close(in)
}
func printer(in chan int) {
	for i := range in {
		fmt.Println("ch ", i)
	}

}
func main() {
	naturals := make(chan int)
	squares := make(chan int)
	//test1(naturals, squares)
	//for i := 0; i < 10; i++ {
	//	fmt.Printf("%d  \n", <-squares)
	//}

	//ch := make(chan int, 3)
	//bufTest(ch)
	//time.Sleep(time.Second * 20)

	go counter(naturals)
	go squarer(naturals, squares)
	printer(squares)
}
