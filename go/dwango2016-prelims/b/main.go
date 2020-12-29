package main

import (
	"bufio"
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
	n := scanInt()
	k := make([]int, n-1)
	ans := make([]int, n)
	for i := range k {
		k[i] = scanInt()
	}

	ans[0], ans[len(ans)-1] = k[0], k[len(k)-1]
	for i := range k[:len(k)-1] {
		if k[i] == k[i+1] {
			ans[i+1] = k[i]
		}
	}
	for i := 1; i < len(ans)-1; i++ {
		if ans[i] == 0 {
			ans[i] = min(k[i-1], k[i])
		}
	}

	wr := bufio.NewWriter(os.Stdout)
	for i := range ans {
		wr.WriteString(strconv.Itoa(ans[i]))
		if i < len(ans)-1 {
			wr.WriteString(" ")
		} else {
			wr.WriteString("\n")
		}
	}
	wr.Flush()
}
