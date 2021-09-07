package main

import (
	"fmt"
	"time"
)

func time1()  {
	now:=time.Now()
	year:=now.Year()
	month:=now.Month()
	day:=now.Day()
	s:= fmt.Sprintf(" %d, %s,%d",year,month,day)
	fmt.Println(s)
	fmt.Println(now.Format(s))
}

//func time2()  {
//
//	fmt.Println()
//}

func main() {
	time1()
	
}
