package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	x := make([]int, n)
	for i := range x {
		sc.Scan()
		x[i], _ = strconv.Atoi(sc.Text())
	}

	sx := make([]int, len(x))
	copy(sx, x)
	sort.Ints(sx)

	lm, rm := sx[len(sx)/2-1], sx[len(sx)/2]
	wr := bufio.NewWriter(os.Stdout)
	for _, v := range x {
		if v <= lm {
			wr.WriteString(strconv.Itoa(rm) + "\n")
		} else {
			wr.WriteString(strconv.Itoa(lm) + "\n")
		}
	}
	wr.Flush()
}
