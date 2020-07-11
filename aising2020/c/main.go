package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	ans := make([]int, n+1)
	idx := 1
	for ; idx*idx <= n; idx++ {
		x := idx
		for y := 1; x*x+y*y <= n; y++ {
			for z := 1; x*x+y*y+z*z <= n; z++ {
				if v := x*x + y*y + z*z + x*y + y*z + z*x; v <= n {
					ans[v]++
				}
			}
		}
	}

	wr := bufio.NewWriter(os.Stdout)
	for _, v := range ans[1:] {
		wr.WriteString(strconv.Itoa(v) + "\n")
	}
	wr.Flush()
}
