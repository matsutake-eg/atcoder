package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	iv, _ := strconv.Atoi(scanString())
	return iv
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 100000), 100000000)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	hs := make([]int, 366)
	for i := range hs {
		if i%7 == 0 || i%7 == 6 {
			hs[i] = 1
		}
	}

	n := scanInt()
	xm := make(map[int]map[int]bool)
	for i := 0; i < n; i++ {
		s := scanString()
		x := strings.Split(s, "/")
		m, _ := strconv.Atoi(x[0])
		d, _ := strconv.Atoi(x[1])
		if _, ok := xm[m]; !ok {
			xm[m] = make(map[int]bool)
		}
		xm[m][d] = true
	}

	t := time.Date(2012, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	for i := range hs {
		m, d := t.Month(), t.Day()
		if xm[int(m)][d] {
			d := i
			for d < len(hs)-1 && hs[d] == 1 {
				d++
			}
			hs[d] = 1
		}
		t = t.AddDate(0, 0, 1)
	}

	ans := 0
	cnt := 0
	for _, v := range hs {
		if v == 1 {
			cnt++
		} else {
			cnt = 0
		}
		ans = max(ans, cnt)
	}
	fmt.Println(ans)
}
