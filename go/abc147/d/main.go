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

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	n := readInt()
	one := make([]int, 60)
	for i := 0; i < n; i++ {
		a := readInt()
		for j := 0; 1<<uint(j) <= a; j++ {
			if b := 1 << uint(j); a&b == b {
				one[j]++
			}
		}
	}

	const mod = 1000000007
	ans := 0
	for i := 0; i < 60; i++ {
		b := 1 << uint(i)
		ans += (one[i] * (n - one[i]) % mod) * (b % mod)
		ans %= mod
	}
	fmt.Println(ans)
}
