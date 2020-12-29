package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	a := make([]int, n)
	for i := range a {
		a[i] = readInt()
	}
	b := make([]int, n)
	for i := range b {
		b[i] = readInt()
	}
	c := make([]int, n)
	for i := range c {
		c[i] = readInt()
	}

	sort.Ints(a)
	sort.Ints(c)
	ans := 0
	for _, v := range b {
		ai := sort.SearchInts(a, v)
		ci := len(c) - sort.SearchInts(c, v+1)
		ans += ai * ci
	}
	fmt.Println(ans)
}
