package main

import (
	"bufio"
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
	h, w, _ := scanInt(), scanInt(), scanInt()
	s := make([]string, h)
	for i := range s {
		s[i] = scanString()
	}

	sumW := make([][]int, h)
	sumH := make([][]int, h)
	for i := range s {
		sumW[i] = make([]int, w)
		sumH[i] = make([]int, w)
		for j := range s[i] {
			if j-1 >= 0 {
				sumW[i][j] = sumW[i][j-1]
			}
			if s[i][j] == '#' {
				sumW[i][j]++
			}
		}
	}
	for j := 0; j < w; j++ {
		for i := 0; i < h; i++ {
			if i-1 >= 0 {
				sumH[i][j] = sumH[i-1][j]
			}
			if s[i][j] == '#' {
				sumH[i][j]++
			}
		}
	}

	ans := make([][]int, h)
	for i := range ans {
		ans[i] = make([]int, w)
	}

	cnt := 1
	li := -1
	for i := range s {
		lj := -1
		for j := range s[i] {
			if s[i][j] == '#' {
				fi := i
				if sumH[i][j] == sumH[h-1][j] {
					fi = h - 1
				}
				fj := j
				if sumW[i][j] == sumW[i][w-1] {
					fj = w - 1
				}
				for ni := li + 1; ni <= fi; ni++ {
					for nj := lj + 1; nj <= fj; nj++ {
						ans[ni][nj] = cnt
					}
				}
				lj = j
				cnt++
			}
		}
		if sumW[i][w-1] > 0 {
			li = i
		}
	}

	var wr = bufio.NewWriter(os.Stdout)
	for _, row := range ans {
		for i, v := range row {
			wr.WriteString(strconv.Itoa(v))
			if i < w-1 {
				wr.WriteString(" ")
			} else {
				wr.WriteString("\n")
			}
		}
	}
	wr.Flush()
}
