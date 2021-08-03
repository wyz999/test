package main

import "fmt"

//两数组的交集
func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, v := range nums1 {
		m[v] += 1
	}
	var nums3 []int
	for _, w := range nums2 {
		if m[w] > 0 {
			nums3 = append(nums3, w)
			delete(m, w)
		}
	}
	return nums3
}

func main() {
	n1 := []int{1, 2, 3, 4, 5, 6}
	n2 := []int{5, 6}
	n3 := intersection(n1, n2)
	fmt.Println(6666)
	fmt.Println(n3)
}
