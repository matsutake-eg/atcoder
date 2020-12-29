package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var a, v, b, w, t int
	fmt.Scan(&a, &v, &b, &w, &t)

	d := abs(a - b)
	if w >= v {
		fmt.Println("NO")
	} else if (v-w)*t >= d {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
