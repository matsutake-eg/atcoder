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
	sc.Buffer(make([]byte, 100000), 100000000)
}

func main() {
	h, w, k := scanInt(), scanInt(), scanInt()
	c := make([]string, h)
	for i := range c {
		c[i] = scanString()
	}

	ans := 0
	for hi := 0; hi < 1<<uint(h); hi++ {
		hm := make(map[int]bool)
		for hj := 0; hj <= h; hj++ {
			if b := 1 << uint(hj); hi&b == b {
				hm[hj] = true
			}
		}

		for wi := 0; wi < 1<<uint(w); wi++ {
			wm := make(map[int]bool)
			for wj := 0; wj < w; wj++ {
				if b := 1 << uint(wj); wi&b == b {
					wm[wj] = true
				}
			}

			cnt := 0
			for ci := range c {
				if hm[ci] {
					continue
				}
				for cj := range c[ci] {
					if wm[cj] {
						continue
					}
					if c[ci][cj] == '#' {
						cnt++
					}
				}
			}
			if cnt == k {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
