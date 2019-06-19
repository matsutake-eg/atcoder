package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	sc := bufio.NewScanner(os.Stdin)
	s := make([]string, h)
	for i := range s {
		sc.Scan()
		s[i] = sc.Text()
	}

	ls := make([][]int, h)
	rs := make([][]int, h)
	for i := range ls {
		ls[i] = make([]int, w)
		for j := 0; j < w; j++ {
			if s[i][j] == '#' {
				ls[i][j] = 0
				continue
			}
			if j == 0 {
				ls[i][j] = 1
				continue
			}
			ls[i][j] = ls[i][j-1] + 1
		}

		rs[i] = make([]int, w)
		for j := w - 1; j >= 0; j-- {
			if s[i][j] == '#' {
				rs[i][j] = 0
				continue
			}
			if j == w-1 {
				rs[i][j] = 1
				continue
			}
			rs[i][j] = rs[i][j+1] + 1
		}
	}

	ts := make([][]int, h)
	ds := make([][]int, h)
	for i := range ts {
		ts[i] = make([]int, w)
		ds[i] = make([]int, w)
	}
	for j := 0; j < w; j++ {
		for i := 0; i < h; i++ {
			if s[i][j] == '#' {
				ts[i][j] = 0
				continue
			}
			if i == 0 {
				ts[i][j] = 1
				continue
			}
			ts[i][j] = ts[i-1][j] + 1
		}

		for i := h - 1; i >= 0; i-- {
			if s[i][j] == '#' {
				ds[i][j] = 0
				continue
			}
			if i == h-1 {
				ds[i][j] = 1
				continue
			}
			ds[i][j] = ds[i+1][j] + 1
		}
	}

	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if v := ls[i][j] + rs[i][j] + ts[i][j] + ds[i][j] - 3; v > ans {
				ans = v
			}
		}
	}
	fmt.Println(ans)
}
