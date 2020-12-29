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

	count := 0
	div_num := 2
	is_div := true

	for {
		for _, v := range nums {
			if v%div_num != 0 {
				is_div = false
				break
			}
		}

		if is_div {
			count++
			div_num *= 2
		} else {
			break
		}
	}

	fmt.Println(count)
}
