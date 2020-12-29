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

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	n := readInt()
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, 10)
		for j := range f[i] {
			f[i][j] = readInt()
		}
	}
	p := make([][]int, n)
	for i := range p {
		p[i] = make([]int, 11)
		for j := range p[i] {
			p[i][j] = readInt()
		}
	}

	ans := -10000000000
	for i := 1; i < 1<<10; i++ {
		o := make([]int, 10)
		for j := range o {
			if b := 1 << uint(j); i&b == b {
				o[j] = 1
			} else {
				o[j] = 0
			}
		}

		sum := 0
		for fi := range f {
			cnt := 0
			for fj := range f[fi] {
				if f[fi][fj] == o[fj] {
					cnt += o[fj]
				}
			}
			sum += p[fi][cnt]
		}
		ans = max(ans, sum)
	}
	fmt.Println(ans)
}
