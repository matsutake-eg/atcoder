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
	n, _ := scanInt(), scanInt()
	xm := make(map[int]int)
	for i := 0; i < n; i++ {
		a := scanInt()
		xm[a]++
		if xm[a]*2 > n {
			fmt.Println(a)
			return
		}
	}
	fmt.Println("?")
}
