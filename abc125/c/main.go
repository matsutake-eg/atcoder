package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func gcd(a, b int) int {
	if b > a {
		a, b = b, a
	}
	// 書き換え可能なパターンの計算簡略化のため
	if b == 0 {
		return a
	}
	var r int
	for {
		r = a % b
		if r == 0 {
			return b
		}
		a = b
		b = r
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		sc.Scan()
		a[i], _ = strconv.Atoi(sc.Text())
	}

	l := make([]int, n+1)
	l[0] = 0
	for i := 0; i < n; i++ {
		l[i+1] = gcd(l[i], a[i])
	}
	r := make([]int, n+1)
	r[n] = 0
	for i := n - 1; i >= 0; i-- {
		r[i] = gcd(r[i+1], a[i])
	}

	max := 0
	for i := 0; i < n; i++ {
		if v := gcd(l[i], r[i+1]); v > max {
			max = v
		}
	}

	fmt.Println(max)
}
