package main

import (
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	iv, _ := strconv.Atoi(readStr())
	return iv
}

func readStr() string {
	sc.Scan()
	return sc.Text()
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	h, w := readInt(), readInt()
	s := make([]string, h)
	for i := range s {
		s[i] = readStr()
	}

	var wr = bufio.NewWriter(os.Stdout)
	mvx := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	mvy := []int{1, 1, 1, 0, 0, -1, -1, -1}
	for i := range s {
		for j := range s[i] {
			if s[i][j] == '#' {
				wr.WriteString("#")
				continue
			}

			cnt := 0
			for k := 0; k < 8; k++ {
				if x, y := j+mvx[k], i+mvy[k]; x >= 0 && x < w && y >= 0 && y < h && s[y][x] == '#' {
					cnt++
				}
			}
			wr.WriteString(strconv.Itoa(cnt))
		}
		wr.WriteString("\n")
	}
	wr.Flush()
}
