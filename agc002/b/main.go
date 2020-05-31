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
	n, m := scanInt(), scanInt()
	a := make([]int, n+1)
	b := make([]bool, n+1)
	for i := range a {
		a[i] = 1
	}
	b[1] = true
	for i := 0; i < m; i++ {
		x, y := scanInt(), scanInt()
		if b[x] {
			if a[x] == 1 {
				b[x] = false
			}
			b[y] = true
		}
		a[x]--
		a[y]++
	}
	ans := 0
	for _, v := range b[1:] {
		if v {
			ans++
		}
	}
	fmt.Println(ans)
}
