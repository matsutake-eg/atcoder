package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	iv, _ := strconv.Atoi(sc.Text())
	return iv
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	h, w := readInt(), readInt()
	c := make([][]int, 10)
	for i := range c {
		c[i] = make([]int, 10)
		for j := range c[i] {
			c[i][j] = readInt()
		}
	}

	for k := 0; k < 10; k++ {
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				c[i][j] = min(c[i][j], c[i][k]+c[k][j])
			}
		}
	}

	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			a := readInt()
			if a == -1 {
				continue
			}
			ans += c[a][1]
		}
	}
	fmt.Println(ans)
}
