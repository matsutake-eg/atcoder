package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	ans := make([]int, 0, 154)
	for i := 0; i <= 153; i++ {
		x := n - i
		if x <= 0 {
			break
		}
		sx := strconv.Itoa(x)
		sum := x
		for _, r := range sx {
			ri, _ := strconv.Atoi(string(r))
			sum += ri
		}
		if sum == n {
			ans = append(ans, x)
		}
	}
	sort.Ints(ans)

	wr := bufio.NewWriter(os.Stdout)
	wr.WriteString(strconv.Itoa(len(ans)) + "\n")
	for _, v := range ans {
		wr.WriteString(strconv.Itoa(v) + "\n")
	}
	wr.Flush()
}
