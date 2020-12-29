package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 100000), 100000000)
}

func main() {
	pa := make([]string, 10)
	for i := range pa {
		pa[i] = scanString()
	}

	type p struct{ i, j int }
	for pi := range pa {
		for pj := range pa[pi] {
			a := make([][]rune, len(pa))
			for i := range a {
				a[i] = make([]rune, len(pa[0]))
				for j := range a[i] {
					a[i][j] = rune(pa[i][j])
				}
			}
			a[pi][pj] = 'o'

			found := false
			seen := make([][]bool, 10)
			for i := range seen {
				seen[i] = make([]bool, 10)
			}
		L:
			for i := range a {
				for j := range a[i] {
					if a[i][j] == 'o' && !seen[i][j] {
						if found {
							break L
						}

						found = true
						ls := list.New()
						ls.PushBack(p{i, j})
						for ls.Len() > 0 {
							v := ls.Remove(ls.Front()).(p)
							seen[v.i][v.j] = true
							if ni := v.i + 1; ni < len(a) && !seen[ni][v.j] && a[ni][v.j] == 'o' {
								ls.PushBack(p{ni, v.j})
							}
							if ni := v.i - 1; ni >= 0 && !seen[ni][v.j] && a[ni][v.j] == 'o' {
								ls.PushBack(p{ni, v.j})
							}
							if nj := v.j + 1; nj < len(a[0]) && !seen[v.i][nj] && a[v.i][nj] == 'o' {
								ls.PushBack(p{v.i, nj})
							}
							if nj := v.j - 1; nj >= 0 && !seen[v.i][nj] && a[v.i][nj] == 'o' {
								ls.PushBack(p{v.i, nj})
							}
						}
					}
					if i == len(a)-1 && j == len(a[0])-1 && found {
						fmt.Println("YES")
						return
					}
				}
			}
		}
	}
	fmt.Println("NO")
}
