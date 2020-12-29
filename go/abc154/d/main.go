package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	iv, _ := strconv.Atoi(sc.Text())
	return iv
}

func max(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	n, k := readInt(), readInt()
	cs := make([]float64, n+1)
	for i := 1; i <= n; i++ {
		p := readInt()
		cs[i] = (float64(p)+1)/2 + cs[i-1]
	}
	ans := float64(0)
	for i := k; i <= n; i++ {
		ans = max(ans, cs[i]-cs[i-k])
	}
	fmt.Println(ans)
}
