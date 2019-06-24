package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	max := (n - 2) * (n - 1) / 2
	if k > max {
		fmt.Println(-1)
		return
	}

	fmt.Println(n - 1 + max - k)

	for i := 1; i <= n-1; i++ {
		fmt.Println(n, i)
	}

	cnt := max
	for i := 1; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if cnt == k {
				return
			}

			cnt--
			fmt.Println(i, j)
		}
	}
}
