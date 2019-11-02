package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	if n == 1 {
		fmt.Println(m)
		return
	}

	ans := 1
	for i := 2; i*i <= m; i++ {
		if m%i != 0 {
			continue
		}
		if i*n <= m && i > ans {
			ans = i
		}
		if j := m / i; j*n <= m && j > ans {
			ans = j
		}
	}
	fmt.Println(ans)
}
