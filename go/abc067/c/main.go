package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	iv, _ := strconv.Atoi(sc.Text())
	return iv
}

func init() {
	sc.Split(bufio.ScanWords)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	n := readInt()
	a := make([]int, n)
	total := 0
	for i := range a {
		a[i] = readInt()
		total += a[i]
	}

	ans := math.MaxInt64
	sum := 0
	for _, v := range a[:len(a)-1] {
		sum += v
		ans = min(ans, abs(total-sum*2))
	}
	fmt.Println(ans)
}
