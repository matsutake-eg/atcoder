package main

import (
	"fmt"
	"sort"
)

func scanNums(nums_len int) (nums []int) {
	var num int
	for i := 0; i < nums_len; i++ {
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	return
}

func main() {
	var n int
	fmt.Scan(&n)
	nums := scanNums(n)

	count := 0
	d := 0
	sort.Sort(sort.IntSlice(nums))
	for _, v := range nums {
		if d < v {
			d = v
			count++
		}
	}

	fmt.Println(count)
}
