package main

import "fmt"

func main() {
	var c1_1, c1_2, c1_3, c2_1, c2_2, c2_3, c3_1, c3_2, c3_3 int
	fmt.Scan(&c1_1, &c1_2, &c1_3, &c2_1, &c2_2, &c2_3, &c3_1, &c3_2, &c3_3)

	d1 := c2_1 - c1_1
	d2 := c2_2 - c1_2
	d3 := c2_3 - c1_3
	if d1 != d2 || d2 != d3 || d3 != d1 {
		fmt.Println("No")
		return
	}
	d1 = c3_1 - c2_1
	d2 = c3_2 - c2_2
	d3 = c3_3 - c2_3
	if d1 != d2 || d2 != d3 || d3 != d1 {
		fmt.Println("No")
		return
	}
	d1 = c1_1 - c3_1
	d2 = c1_2 - c3_2
	d3 = c1_3 - c3_3
	if d1 != d2 || d2 != d3 || d3 != d1 {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
}
