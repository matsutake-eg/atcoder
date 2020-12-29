package main

import "fmt"

func getWin(h rune) rune {
	if h == 'r' {
		return 'p'
	}
	if h == 's' {
		return 'r'
	}
	return 's'
}

func main() {
	var (
		n, k, r, s, p int
		t             string
	)
	fmt.Scan(&n, &k, &r, &s, &p, &t)

	hs := []rune{'r', 's', 'p'}
	ans := 0
	for i := 0; i < k; i++ {
		bh := 'x'
		for j := i; j < n; j += k {
			h := getWin(rune(t[j]))
			if h == bh {
				if j+k >= n {
					break
				}

				nh := getWin(rune(t[j+k]))
				for _, v := range hs {
					if v == bh || v == nh {
						continue
					}
					h = v
				}
			} else {
				switch h {
				case 'r':
					ans += r
				case 's':
					ans += s
				case 'p':
					ans += p
				}
			}
			bh = h
		}
	}
	fmt.Println(ans)
}
