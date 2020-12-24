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
	n, c, k := scanInt(), scanInt(), scanInt()
	t := make([]int, n)
	for i := range t {
		t[i] = scanInt()
	}
	sort.Ints(t)

	ans := 1
	l := t[0] + k
	x := 1
	for i := 1; i < n; i++ {
		x++
		if x > c || t[i] > l {
			ans++
			l = t[i] + k
			x = 1
		}
	}
	fmt.Println(ans)
}
