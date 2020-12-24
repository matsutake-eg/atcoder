package main

import "fmt"

func scanNums(n int, s string) [][]int {
	nums := make([][]int, 2)
	for i := 0; i < 2; i++ {
		nums[i] = make([]int, n)
	}

	nums[0][0] = 0
	nums[1][n-1] = 0
	for i := 1; i < n; i++ {
		nums[0][i] = nums[0][i-1] + count(s[i-1], 'W')
	}
	for i := n - 2; i >= 0; i-- {
		nums[1][i] = nums[1][i+1] + count(s[i+1], 'E')
	}

	return nums
}

func count(b byte, r byte) int {
	if b == r {
		return 1
	}
	return 0
}

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)
	nums := scanNums(n, s)

	min := n
	for i := 0; i < n; i++ {
		if nums[0][i]+nums[1][i] < min {
			min = nums[0][i] + nums[1][i]
		}
	}

	fmt.Println(min)
}
