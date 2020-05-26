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

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	k, t := scanInt(), scanInt()
	mx := math.MinInt64
	for i := 0; i < t; i++ {
		a := scanInt()
		mx = max(mx, a)
	}
	fmt.Println(max(mx-1-(k-mx), 0))
}
