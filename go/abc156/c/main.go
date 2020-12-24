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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	n := readInt()
	x := make([]int, n)
	for i := range x {
		x[i] = readInt()
	}

	ans := math.MaxInt64
	for i := 1; i <= 100; i++ {
		sum := 0
		for _, v := range x {
			sum += (v - i) * (v - i)
		}
		ans = min(ans, sum)
	}
	fmt.Println(ans)
}
