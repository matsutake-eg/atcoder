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
	na, nb := scanInt(), scanInt()
	xm := make(map[int]bool)
	for i := 0; i < na; i++ {
		a := scanInt()
		xm[a] = true
	}
	cnt := 0
	for i := 0; i < nb; i++ {
		b := scanInt()
		if xm[b] {
			cnt++
		} else {
			xm[b] = true
		}
	}
	fmt.Println(float64(cnt) / float64(len(xm)))
}
