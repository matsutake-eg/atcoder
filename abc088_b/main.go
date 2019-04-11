package main

import (
	"fmt"
	"sort"
)

func scanNums(len int) (nums []int) {
	var num int
	for i := 0; i < len; i++ {
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	return
}

func main() {
	var n int
	fmt.Scan(&n)
	nums := scanNums(n)
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	i := 0
	alice := 0
	bob := 0
	nums_len := len(nums)
	for {
		if i < nums_len {
			alice += nums[i]
		} else {
			break
		}
		i++

		if i < nums_len {
			bob += nums[i]
		} else {
			break
		}
		i++
	}

	fmt.Println(alice - bob)
}
