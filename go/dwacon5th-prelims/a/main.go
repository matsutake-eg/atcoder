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

func main() {
	n := scanInt()
	ave := 0.0
	type flame struct{ a, n int }
	fs := make([]flame, n)
	for i := range fs {
		fs[i].a, fs[i].n = scanInt(), i
		ave += float64(fs[i].a)
	}
	ave /= float64(n)

	d := math.MaxFloat64
	var ans int
	for _, f := range fs {
		if v := abs(ave - float64(f.a)); v < d {
			d = v
			ans = f.n
		}
	}
	fmt.Println(ans)
}
