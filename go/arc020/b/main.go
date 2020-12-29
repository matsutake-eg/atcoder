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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	n, c := scanInt(), scanInt()
	xm := make(map[int]int)
	ym := make(map[int]int)
	for i := 0; i < n; i++ {
		a := scanInt()
		if i%2 == 0 {
			xm[a]++
		} else {
			ym[a]++
		}
	}

	sum := math.MaxInt64
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			if i == j {
				continue
			}
			sum = min(sum, n-xm[i]-ym[j])
		}
	}
	fmt.Println(sum * c)
}
