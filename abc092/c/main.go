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
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	a := make([]int, n+2)
	a[0], a[len(a)-1] = 0, 0
	for i := 1; i < len(a)-1; i++ {
		sc.Scan()
		a[i], _ = strconv.Atoi(sc.Text())
	}
	mv := 0
	for i := 1; i < len(a); i++ {
		mv += abs(a[i-1] - a[i])
	}

	wr := bufio.NewWriter(os.Stdout)
	for i := 1; i < len(a)-1; i++ {
		ans := mv - (abs(a[i-1]-a[i]) + abs(a[i+1]-a[i])) + abs(a[i+1]-a[i-1])
		wr.WriteString(strconv.Itoa(ans) + "\n")
	}
	wr.Flush()
}
