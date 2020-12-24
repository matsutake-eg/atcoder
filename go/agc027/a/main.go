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
	n, x := scanInt(), scanInt()
	a := make([]int, n)
	for i := range a {
		a[i] = scanInt()
	}
	sort.Ints(a)

	for i := range a {
		x -= a[i]
		if x < 0 {
			fmt.Println(i)
			return
		}
	}
	if x == 0 {
		fmt.Println(n)
	} else {
		fmt.Println(n - 1)
	}
}
