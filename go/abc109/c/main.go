package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(a, b int) int {
	r := a % b
	if r == 0 {
		return b
	}
	return gcd(b, r)
}

func gcdMulti(xs []int) int {
	ans := xs[0]
	for i := 1; i < len(xs); i++ {
		ans = gcd(ans, xs[i])
	}
	return ans
}

func main() {
	var N, X int
	fmt.Scan(&N, &X)
	ds := make([]int, N)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for i := range ds {
		sc.Scan()
		t, _ := strconv.Atoi(sc.Text())
		ds[i] = abs(X - t)
	}
	fmt.Println(gcdMulti(ds))
}
