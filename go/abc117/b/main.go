package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	l := make([]int, n)
	for i := range l {
		fmt.Scan(&l[i])
	}

	for i, v := range l {
		left := 0
		for j := 0; j < i; j++ {
			left += l[j]
		}
		right := 0
		for j := i + 1; j < n; j++ {
			right += l[j]
		}
		if v >= left+right {
			fmt.Println("No")
			break
		}
		if i == n-1 {
			fmt.Println("Yes")
		}
	}
}
