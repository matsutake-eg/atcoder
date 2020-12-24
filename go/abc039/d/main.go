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
	s := make([]string, h)
	for i := 0; i < h; i++ {
		s[i] = scanString()
	}

	dx := []int{0, 1, 1, 1, 0, -1, -1, -1}
	dy := []int{1, 1, 0, -1, -1, -1, 0, 1}
	ans := make([][]byte, h)
	for i := 0; i < h; i++ {
		ans[i] = make([]byte, w)
		for j := 0; j < w; j++ {
			ans[i][j] = s[i][j]
			if ans[i][j] == '.' {
				continue
			}

			for k := 0; k < 8; k++ {
				nx, ny := j+dx[k], i+dy[k]
				if nx < 0 || nx >= w || ny < 0 || ny >= h || s[ny][nx] == '#' {
					continue
				}

				ans[i][j] = '.'
				break
			}
		}
	}

	ns := make([][]byte, h)
	for i := 0; i < h; i++ {
		ns[i] = make([]byte, w)
		for j := 0; j < w; j++ {
			ns[i][j] = ans[i][j]
			for k := 0; k < 8; k++ {
				nx, ny := j+dx[k], i+dy[k]
				if nx < 0 || nx >= w || ny < 0 || ny >= h {
					continue
				}

				if ans[ny][nx] == '#' {
					ns[i][j] = '#'
					break
				}
			}
		}
	}

	for i := range ns {
		if string(ns[i]) != s[i] {
			fmt.Println("impossible")
			return
		}
	}
	fmt.Println("possible")
	for _, v := range ans {
		fmt.Println(string(v))
	}
}
