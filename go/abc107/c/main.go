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

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	x := make([]int, n)
	for i := range x {
		sc.Scan()
		x[i], _ = strconv.Atoi(sc.Text())
	}

	k--
	ans := 1000000000
	for i := 0; i+k < n; i++ {
		d := x[i+k] - x[i]
		if v := abs(x[i]) + d; v < ans {
			ans = v
		}
		if v := abs(x[i+k]) + d; v < ans {
			ans = v
		}
	}
	fmt.Println(ans)
}
