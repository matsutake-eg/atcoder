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
	xm := make(map[int]bool)
	for i := 0; i < 6; i++ {
		e := scanInt()
		xm[e] = true
	}
	b := scanInt()

	isb := false
	cnt := 0
	for i := 0; i < 6; i++ {
		l := scanInt()
		if l == b {
			isb = true
		}
		if xm[l] {
			cnt++
		}
	}

	switch cnt {
	case 6:
		fmt.Println(1)
	case 5:
		if isb {
			fmt.Println(2)
		} else {
			fmt.Println(3)
		}
	case 4:
		fmt.Println(4)
	case 3:
		fmt.Println(5)
	default:
		fmt.Println(0)
	}
}
