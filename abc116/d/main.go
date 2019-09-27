package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type sushi struct{ t, d int }
type sushis []sushi

func (s sushis) Len() int           { return len(s) }
func (s sushis) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sushis) Less(i, j int) bool { return s[i].d > s[j].d }

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	s := sushis(make([]sushi, n))
	for i := range s {
		sc.Scan()
		s[i].t, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		s[i].d, _ = strconv.Atoi(sc.Text())
	}

	sort.Sort(s)
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
