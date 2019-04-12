package main

import "fmt"

func scanMultiNums(multi_nums_len int) (multi_nums [][]int) {
	multi_nums = append(multi_nums, []int{0, 0, 0})

	var t, x, y int
	for i := 0; i < multi_nums_len; i++ {
		fmt.Scan(&t, &x, &y)
		multi_nums = append(multi_nums, []int{t, x, y})
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var n int
	fmt.Scan(&n)
	multi_nums := scanMultiNums(n)

	var is_travel bool
	var dt, dx, dy int

	for i := 1; i < n+1; i++ {
		dt = multi_nums[i][0] - multi_nums[i-1][0]
		dx = abs(multi_nums[i][1] - multi_nums[i-1][1])
		dy = abs(multi_nums[i][2] - multi_nums[i-1][2])

		if dt < dx || dt < dy {
			is_travel = false
			break
		} else if dt%2 != (dx+dy)%2 {
			is_travel = false
			break
		} else {
			is_travel = true
		}
	}

	if is_travel {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
