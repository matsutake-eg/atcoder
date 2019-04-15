package main

import "fmt"

func scanStrs(h int) (strs [][]rune) {
	strs = make([][]rune, h)
	var s string
	for i := range strs {
		fmt.Scan(&s)
		strs[i] = []rune(s)
	}
	return
}

func drawable(strs [][]rune, i, j int) bool {
	if strs[i][j] == '.' {
		return true
	}

	result := false
	if i-1 >= 0 {
		if strs[i-1][j] == '#' {
			result = true
		}
	}

	if j-1 >= 0 {
		if strs[i][j-1] == '#' {
			result = true
		}
	}

	if j+1 < len(strs[0]) {
		if strs[i][j+1] == '#' {
			result = true
		}
	}

	if i+1 < len(strs) {
		if strs[i+1][j] == '#' {
			result = true
		}
	}

	return result
}

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	strs := scanStrs(h)

	is_complete := true
	for i := 0; i < h; i++ {
		if !is_complete {
			break
		}
		for j := 0; j < w; j++ {
			if !is_complete {
				break
			}
			is_complete = drawable(strs, i, j)
		}
	}

	if is_complete {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
