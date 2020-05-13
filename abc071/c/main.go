package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	iv, _ := strconv.Atoi(sc.Text())
	return iv
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	n := readInt()
	m := make(map[int]int, n)
	for i := 0; i < n; i++ {
		a := readInt()
		m[a]++
	}

	type stick struct{ a, c int }
	sts := make([]stick, 0, len(m))
	for a, c := range m {
		sts = append(sts, stick{a, c})
	}
	sort.Slice(sts, func(i, j int) bool { return sts[i].a > sts[j].a })
	t := 0
	for _, st := range sts {
		if t == 0 {
			if st.c >= 4 {
				fmt.Println(st.a * st.a)
				return
			} else if st.c >= 2 {
				t = st.a
			}
		} else {
			if st.c >= 2 {
				fmt.Println(t * st.a)
				return
			}
		}
	}
	fmt.Println(0)
}
