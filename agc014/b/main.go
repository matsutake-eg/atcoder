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
	_, m := scanInt(), scanInt()
	xm := make(map[int]int)
	for i := 0; i < m; i++ {
		a, b := scanInt(), scanInt()
		xm[a]++
		xm[b]++
	}

	for _, v := range xm {
		if v%2 == 1 {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
