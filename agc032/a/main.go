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
	n := scanInt()
	b := make([]int, n)
	for i := range b {
		b[i] = scanInt()
	}

	ans := make([]int, 0, n)
	for {
		if len(b) == 0 {
			break
		}
		for i := len(b) - 1; i >= 0; i-- {
			if b[i] == i+1 {
				ans = append(ans, b[i])
				b = append(b[:i], b[i+1:]...)
				break
			}

			if i == 0 {
				fmt.Println(-1)
				return
			}
		}
	}

	wr := bufio.NewWriter(os.Stdout)
	for i := n - 1; i >= 0; i-- {
		wr.WriteString(strconv.Itoa(ans[i]) + "\n")
	}
	wr.Flush()
}
