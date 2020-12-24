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
	a := make([]int, n)
	for i := range a {
		a[i] = scanInt()
	}
	b := make([]int, n)
	for i := range b {
		b[i] = scanInt()
	}

	c := make([]int, n)
	for i := range c {
		c[i] = a[i] - b[i]
	}
	sort.Ints(c)

	ans := 0
	sum := 0
	for _, v := range c {
		if v >= 0 {
			break
		}
		ans++
		sum += v
	}
	for i := n - 1; i >= 0; i-- {
		if c[i] <= 0 || sum >= 0 {
			break
		}
		ans++
		sum += c[i]
	}

	if sum < 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}
