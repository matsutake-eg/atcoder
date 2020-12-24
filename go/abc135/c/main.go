package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	a := make([]int, n+1)
	for i := range a {
		sc.Scan()
		a[i], _ = strconv.Atoi(sc.Text())
	}

	ans := 0
	for i := 0; i < n; i++ {
		sc.Scan()
		b, _ := strconv.Atoi(sc.Text())
		if b <= a[i] {
			ans += b
			continue
		}
		ans += a[i]
		b = b - a[i]

		if b <= a[i+1] {
			a[i+1] -= b
			ans += b
			continue
		}
		ans += a[i+1]
		a[i+1] = 0
	}
	fmt.Println(ans)
}
