package main

import "fmt"

func scanNums(l int) [][]int {
	nums := make([][]int, 2)
	var n, sum int
	for i := 0; i < 2; i++ {
		nums[i] = make([]int, l)
		sum = 0
		for j := 0; j < l; j++ {
			fmt.Scan(&n)
			sum += n
			nums[i][j] = sum
		}
	}
	return nums
}

func main() {
	var n int
	fmt.Scan(&n)
	nums := scanNums(n)

	max := nums[0][0] + nums[1][n-1]
	var total int
	for i := 1; i < n; i++ {
		total = nums[0][i] + nums[1][n-1] - nums[1][i-1]
		if total > max {
			max = total
		}
	}

	fmt.Println(max)
}
