package main

import (
	"bufio"
	"fmt"
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
	xs := make([][]int, 4)
	for i := range xs {
		xs[i] = make([]int, 4)
		for j := range xs[i] {
			xs[i][j] = scanInt()
		}
	}

	for i := range xs {
		for j := range xs[i] {
			pi := i - 1
			ni := i + 1
			pj := j - 1
			nj := j + 1
			if pi >= 0 && xs[pi][j] == xs[i][j] {
				fmt.Println("CONTINUE")
				return
			} else if pj >= 0 && xs[i][pj] == xs[i][j] {
				fmt.Println("CONTINUE")
				return
			} else if ni < 4 && xs[ni][j] == xs[i][j] {
				fmt.Println("CONTINUE")
				return
			} else if nj < 4 && xs[i][nj] == xs[i][j] {
				fmt.Println("CONTINUE")
				return
			}
		}
	}
	fmt.Println("GAMEOVER")
}
