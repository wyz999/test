package main

import (
	"fmt"
)

func gotest(n int) int {
	a, b := 0, 1
	for n > 0 {
		a, b = b, a+b
		fmt.Println(a)
		n--
	}
	return a
}

func main() {
	fmt.Println("666666")
	fmt.Println("git test")
	fmt.Println(gotest(6))

}
