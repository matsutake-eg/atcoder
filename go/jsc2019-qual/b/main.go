package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const mod = 1000000007

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	a := make([]int, n)
	for i := range a {
		sc.Scan()
		a[i], _ = strconv.Atoi(sc.Text())
	}

	ac, cc := 0, 0
	for i := range a {
		for j := range a {
			if a[i] <= a[j] {
				continue
			}
			ac++
			if j >= i {
				cc++
			}
		}
	}
	ans := k * (k - 1) / 2 % mod
	ans = (ans*ac + k*cc) % mod
	fmt.Println(ans)
}
