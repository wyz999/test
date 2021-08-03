package main

import "fmt"

type Linknode struct {
	Data int
	Next *Linknode
}

func main() {
	l1 := new(Linknode)
	l1.Data = 123
	l2 := new(Linknode)
	l2.Data = 456
	l3 := new(Linknode)
	l3.Data = 789
	l1.Next = l2
	l2.Next = l3
	nodeNew := l1
	for {
		if nodeNew != nil {
			fmt.Println(nodeNew.Data)
			nodeNew = nodeNew.Next
			continue
		}
		break
	}

}
