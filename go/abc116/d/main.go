package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	n, k := scanInt(), scanInt()
	type sushi struct{ t, d int }
	s := make([]sushi, n)
	for i := range s {
		s[i].t, s[i].d = scanInt(), scanInt()
	}

	sort.Slice(s, func(i, j int) bool { return s[i].d > s[j].d })
	sumD := 0
	mt := make(map[int]bool, n)
	src := make([]sushi, 0, k)
	for i := 0; i < k; i++ {
		sumD += s[i].d
		if mt[s[i].t] {
			src = append(src, s[i])
			continue
		}
		mt[s[i].t] = true
	}
	ts := len(mt)

	dst := make([]sushi, 0, len(s)-k+1)
	for i := k; i < len(s); i++ {
		if !mt[s[i].t] {
			mt[s[i].t] = true
			dst = append(dst, s[i])
		}
	}

	ans := sumD + ts*ts
	for i := range src {
		if i >= len(dst) {
			break
		}
		sumD += -src[len(src)-i-1].d + dst[i].d
		ts++
		if v := sumD + ts*ts; v > ans {
			ans = v
		}
	}
	fmt.Println(ans)
}
