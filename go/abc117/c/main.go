package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	x := make([]int, m)
	for i := range x {
		sc.Scan()
		x[i], _ = strconv.Atoi(sc.Text())
	}

	sort.Ints(x)
	d := make([]int, m-1)
	for i := range d {
		d[i] = x[i+1] - x[i]
	}
	sort.Ints(d)
	sum := 0
	for i := 0; i < n-1; i++ {
		if v := m - 2 - i; v < 0 {
			break
		} else {
			sum += d[v]
		}
	}
	fmt.Println(x[m-1] - x[0] - sum)
}
