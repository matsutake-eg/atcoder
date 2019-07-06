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

	sum := 0
	min := 10000000001
	isNegaOdd := false
	for i := 0; i < n; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		if a < 0 {
			isNegaOdd = !isNegaOdd
			a = -a
		}
		sum += a
		if a < min {
			min = a
		}
	}
	if isNegaOdd {
		sum -= min * 2
	}
	fmt.Println(sum)
}
