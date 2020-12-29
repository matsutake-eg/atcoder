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

func mins(x ...int) int {
	mv := x[0]
	for _, v := range x[1:] {
		if v < mv {
			mv = v
		}
	}
	return mv
}

func main() {
	n, a, b := scanInt(), scanInt(), scanInt()
	ca, cb, cc := 0, 0, 0
	for i := 0; i < n; i++ {
		p := scanInt()
		if p <= a {
			ca++
		} else if p >= b+1 {
			cc++
		} else {
			cb++
		}
	}
	fmt.Println(mins(ca, cb, cc))
}
