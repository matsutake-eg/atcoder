package main

import (
	"bufio"
	"fmt"
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

func main() {
	n, k := scanInt(), scanInt()
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = scanInt()
	}

	ps := make([]int, n+1)
	for i := 2; i <= n; i++ {
		ps[i] = -1
	}

	ans := make([]int, n+1)
	ans[0] = 1

	var b, r int
	p := 1
	for i := 1; i <= n; i++ {
		p = a[p]
		if i == k {
			fmt.Println(p)
			return
		}

		if ps[p] != -1 {
			b = ps[p]
			r = i - b
			break
		}
		ps[p] = i
		ans[i] = p
	}

	k -= b
	k %= r
	fmt.Println(ans[b+k])
}
