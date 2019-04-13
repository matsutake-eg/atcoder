package main

import "fmt"

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
	nums = append(nums, 101)

	sum := 0
	max := 0
	for i := 0; i < n; i++ {
		if i == 0 {
			sum++
		} else if max <= nums[i] {
			sum++
		}

		if nums[i] > max {
			max = nums[i]
		}
	}

	fmt.Println(sum)
}
