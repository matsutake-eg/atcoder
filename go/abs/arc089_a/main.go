package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var n int
	fmt.Scan(&n)

	var is_travel bool
	var t, x, y, t_b, x_b, y_b, dt, dx, dy int

	for i := 0; i < n; i++ {
		fmt.Scan(&t, &x, &y)
		dt = t - t_b
		dx = abs(x - x_b)
		dy = abs(y - y_b)

		if dt < dx || dt < dy {
			is_travel = false
		} else if dt%2 != (dx+dy)%2 {
			is_travel = false
		} else {
			is_travel = true
		}

		if is_travel {
			t_b = t
			x_b = x
			y_b = y
		} else {
			break
		}
	}

	if is_travel {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
