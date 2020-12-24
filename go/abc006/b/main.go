package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	switch n {
	case 1:
		fmt.Println(0)
		return
	case 2:
		fmt.Println(0)
		return
	case 3:
		fmt.Println(1)
		return
	case 4:
		fmt.Println(1)
		return
	}

	ans := make([]int, n+1)
	ans[1], ans[2], ans[3], ans[4] = 0, 0, 1, 1
	sum := 1
	for i := 5; i <= n; i++ {
		sum -= ans[i-4]
		sum += ans[i-1]
		ans[i] = sum % 10007
	}
	fmt.Println(ans[n])
}
