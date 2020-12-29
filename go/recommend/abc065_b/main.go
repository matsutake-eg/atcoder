package main

import "fmt"

func scanNums(nums_len int) []int {
	nums := make([]int, nums_len)
	var num int
	for i := 0; i < nums_len; i++ {
		fmt.Scan(&num)
		nums[i] = num
	}
	return nums
}

func main() {
	var n int
	fmt.Scan(&n)
	nums := scanNums(n)

	count := 0
	pos := 1
	for i := 0; i < n; i++ {
		pos = nums[pos-1]
		count++

		if pos == 2 {
			break
		}
	}

	if count == n {
		fmt.Println(-1)
	} else {
		fmt.Println(count)
	}
}
