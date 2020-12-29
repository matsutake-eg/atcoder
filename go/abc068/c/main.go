package main

import (
	"bufio"
	"fmt"
	"os"
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
	n, m := readInt(), readInt()
	xm := make(map[int]map[int]bool, n)
	for i := 0; i < m; i++ {
		a, b := readInt(), readInt()
		if _, ok := xm[a]; !ok {
			xm[a] = make(map[int]bool)
		}
		xm[a][b] = true
	}

	for t := range xm[1] {
		for t2 := range xm[t] {
			if t2 == n {
				fmt.Println("POSSIBLE")
				return
			}
		}
	}
	fmt.Println("IMPOSSIBLE")
}
