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
	sort.Ints(a)

	sum := 0
	ans := 0
	for i := 0; i < n-1; i++ {
		ans++
		sum += a[i]
		if sum*2 < a[i+1] {
			ans = 0
		}
	}
	fmt.Println(ans + 1)
}
