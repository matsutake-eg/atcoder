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

func main() {
	n := scanInt()
	var (
		mx  = math.MinInt64
		ans int
	)
	for i := 0; i < n; i++ {
		a, b := scanInt(), scanInt()
		if a > mx {
			mx = a
			ans = a + b
		}
	}
	fmt.Println(ans)
}
