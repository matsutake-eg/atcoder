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
	var n, t, a int
	fmt.Scan(&n, &t, &a)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	var x int
	min := 1000000000
	minIndex := 0
	for i := 0; i < n; i++ {
		sc.Scan()
		x, _ = strconv.Atoi(sc.Text())
		if v := abs(a*1000 - t*1000 + x*6); v < min {
			min = v
			minIndex = i + 1
		}
	}
	fmt.Println(minIndex)
}
