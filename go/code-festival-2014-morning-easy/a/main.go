package main

import (
	"bufio"
	"fmt"
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
	a := make([]float64, n)
	for i := range a {
		a[i] = float64(scanInt())
	}

	sum := 0.0
	for i := 1; i < n; i++ {
		sum += a[i] - a[i-1]
	}
	fmt.Println(sum / float64(n-1))
}
