package main

import (
	"bufio"
	"fmt"
	"math"
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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	n := scanInt()
	a := make([]int, n)
	for i := range a {
		a[i] = scanInt()
	}

	ans := math.MaxInt64
	for i := -100; i <= 100; i++ {
		sum := 0
		for _, v := range a {
			sum += (v - i) * (v - i)
		}
		ans = min(ans, sum)
	}
	fmt.Println(ans)
}
