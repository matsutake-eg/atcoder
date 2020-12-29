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
	xm := make(map[int]bool)
	for i := 0; i < k; i++ {
		d := scanInt()
		for i := 0; i < d; i++ {
			a := scanInt()
			xm[a] = true
		}
	}
	fmt.Println(n - len(xm))
}
