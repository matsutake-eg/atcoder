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

func max(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

func main() {
	n := scanInt()
	xs := make([]float64, n)
	ys := make([]float64, n)
	for i := 0; i < n; i++ {
		xs[i], ys[i] = float64(scanInt()), float64(scanInt())
	}
	ans := 0.0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			d := math.Sqrt(math.Pow(xs[i]-xs[j], 2) + math.Pow(ys[i]-ys[j], 2))
			ans = max(ans, d)
		}
	}
	fmt.Println(ans)
}
