package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func max(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

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
	n := scanInt()
	ans := math.SmallestNonzeroFloat64
	for i := 0; i < n; i++ {
		a, b, c, d, e := scanInt(), scanInt(), scanInt(), scanInt(), scanInt()
		ans = max(ans, float64(a+b+c+d)+float64(e*110)/900)
	}
	fmt.Println(ans)
}
