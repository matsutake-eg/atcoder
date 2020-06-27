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

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func main() {
	x, y := float64(scanInt()), float64(scanInt())
	n := scanInt()
	xs := make([]float64, n+1)
	ys := make([]float64, n+1)
	for i := 0; i < n; i++ {
		xs[i], ys[i] = float64(scanInt()), float64(scanInt())
	}
	xs[n], ys[n] = xs[0], ys[0]

	ans := math.MaxFloat64
	for i := 1; i <= n; i++ {
		dx, dy := xs[i-1]-xs[i], ys[i-1]-ys[i]
		ans = min(ans, abs(dx*(ys[i]-y)-dy*(xs[i]-x))/math.Sqrt(dx*dx+dy*dy))
	}
	fmt.Println(ans)
}
