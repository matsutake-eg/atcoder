package main

import (
	"bufio"
	"fmt"
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
	t := make([]int, n)
	tm := make(map[int]int)
	for i := range t {
		t[i] = scanInt()
		tm[t[i]]++
	}
	sort.Ints(t)

	sum := 0
	for i, v := range t {
		sum += v * (len(t) - i)
	}
	fmt.Println(sum)

	ans := 1
	const mod = 1000000007
	for _, v := range tm {
		if v == 1 {
			continue
		}
		for i := 2; i <= v; i++ {
			ans *= i
			ans %= mod
		}
	}
	fmt.Println(ans)
}
