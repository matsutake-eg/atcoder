package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	sc.Buffer(make([]byte, 100001), 100001*100)
}

func main() {
	h, w := scanInt(), scanInt()
	xm := make(map[rune]int)
	for i := 0; i < h; i++ {
		a := scanString()
		for _, v := range a {
			xm[v]++
		}
	}

	var g1, g2, g4 int
	if h%2 == 1 && w%2 == 1 {
		g1 = 1
	}
	if h%2 == 1 {
		g2 += w / 2
	}
	if w%2 == 1 {
		g2 += h / 2
	}
	g4 = (h / 2) * (w / 2)

	for i := 0; i < g4; i++ {
		found := false
		for j := range xm {
			if xm[j] >= 4 {
				xm[j] -= 4
				found = true
				break
			}
		}
		if !found {
			fmt.Println("No")
			return
		}
	}
	for i := 0; i < g2; i++ {
		found := false
		for j := range xm {
			if xm[j] >= 2 {
				xm[j] -= 2
				found = true
				break
			}
		}
		if !found {
			fmt.Println("No")
			return
		}
	}
	if g1 == 1 {
		found := false
		for j := range xm {
			if xm[j] >= 1 {
				xm[j]--
				found = true
				break
			}
		}
		if !found {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
