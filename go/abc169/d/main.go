package main

import "fmt"

func prmFac(x int) map[int]bool {
	pfm := make(map[int]bool)
	for i := 2; i*i <= x; i++ {
		for x%i == 0 {
			x /= i
			pfm[i] = true
		}
	}
	if x > 1 {
		pfm[x] = true
	}
	return pfm
}

func main() {
	var n int
	fmt.Scan(&n)

	ans := 0
	for i := range prmFac(n) {
		t := 1
		for {
			t *= i
			if n%t != 0 {
				break
			}
			n /= t
			ans++
		}
	}
	fmt.Println(ans)
}
