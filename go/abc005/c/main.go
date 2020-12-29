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
	t, n := scanInt(), scanInt()
	ts := make([]int, 101)
	for i := 0; i < n; i++ {
		a := scanInt()
		ts[a]++
	}

	m := scanInt()
	cs := make([]int, 101)
	for i := 0; i < m; i++ {
		b := scanInt()
		cs[b]++
	}

	for i, c := range cs {
		if c > 0 {
			cnt := 0
			for j := i - t; j <= i; j++ {
				if j < 0 || j >= len(cs) {
					continue
				}
				if ts[j] > 0 {
					if cnt+ts[j] > c {
						ts[j] -= c - cnt
						cnt = c
						break
					} else {
						cnt += ts[j]
						ts[j] = 0
					}
				}
			}
			if cnt < c {
				fmt.Println("no")
				return
			}
		}
	}
	fmt.Println("yes")
}
