package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

func main() {
	n := scanInt()
	r := make([]int, n)
	for i := range r {
		r[i] = scanInt()
	}

	sort.Sort(sort.Reverse(sort.IntSlice(r)))
	ans := 0
	for i := 0; i < n; i += 2 {
		ans += r[i] * r[i]
		if i+1 < n {
			ans -= r[i+1] * r[i+1]
		}
	}
	fmt.Println(float64(ans) * math.Pi)
}
