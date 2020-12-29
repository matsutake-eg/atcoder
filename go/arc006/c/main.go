package main

import (
	"bufio"
	"fmt"
	"math"
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
	n := scanInt()
	xm := make(map[int]bool)
	ans := 0
	for i := 0; i < n; i++ {
		w := scanInt()
		t := -1
		d := math.MaxInt64
		for i := range xm {
			if v := i - w; v >= 0 && v < d {
				d = v
				t = i
			}
		}
		if t == -1 {
			ans++
		} else {
			delete(xm, t)
		}
		xm[w] = true
	}
	fmt.Println(ans)
}
