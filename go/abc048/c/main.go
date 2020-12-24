package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	iv, _ := strconv.Atoi(sc.Text())
	return iv
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	n, x := scanInt(), scanInt()
	ans := 0
	a := make([]int, n)
	for i := range a {
		a[i] = scanInt()
		if d := a[i] - x; d > 0 {
			ans += d
			a[i] = x
		}
	}

	for i := range a[:n-1] {
		if d := a[i] + a[i+1] - x; d > 0 {
			ans += d
			a[i+1] -= d
		}
	}
	fmt.Println(ans)
}
