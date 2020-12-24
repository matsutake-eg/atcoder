package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	cnt := make([]int, 5)
	for i := 0; i < n; i++ {
		sc.Scan()
		s := sc.Text()
		switch s[0] {
		case 'M':
			cnt[0]++
		case 'A':
			cnt[1]++
		case 'R':
			cnt[2]++
		case 'C':
			cnt[3]++
		case 'H':
			cnt[4]++
		}
	}
	p := []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 2}
	q := []int{1, 1, 1, 2, 2, 3, 2, 2, 3, 3}
	r := []int{2, 3, 4, 3, 4, 4, 3, 4, 4, 4}
	ans := 0
	for i := 0; i < 10; i++ {
		ans += cnt[p[i]] * cnt[q[i]] * cnt[r[i]]
	}
	fmt.Println(ans)
}
