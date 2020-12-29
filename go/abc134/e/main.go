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

	sc := bufio.NewScanner(os.Stdin)
	ans := make([]int, 0, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())

		if len(ans) == 0 || a <= ans[len(ans)-1] {
			ans = append(ans, a)
			continue
		}

		p := sort.Search(len(ans), func(i int) bool { return ans[i] < a })
		ans[p] = a
	}
	fmt.Println(len(ans))
}
