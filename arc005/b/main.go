package main

import (
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	iv, _ := strconv.Atoi(scanString())
	return iv
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 100001), 100001*100)
}

func main() {
	x, y, w := scanInt(), scanInt(), scanString()
	x--
	y--
	c := make([]string, 9)
	for i := range c {
		c[i] = scanString()
	}

	type p struct{ dx, dy int }
	mv := make(map[string]p)
	mv["R"] = p{1, 0}
	mv["L"] = p{-1, 0}
	mv["U"] = p{0, -1}
	mv["D"] = p{0, 1}
	mv["RU"] = p{1, -1}
	mv["RD"] = p{1, 1}
	mv["LU"] = p{-1, -1}
	mv["LD"] = p{-1, 1}

	wr := bufio.NewWriter(os.Stdout)
	for i := 0; i < 4; i++ {
		wr.WriteString(string(c[y][x]))
		dx, dy := mv[w].dx, mv[w].dy
		switch w {
		case "R":
			if x == 8 {
				mv[w] = p{-dx, dy}
			}
		case "L":
			if x == 0 {
				mv[w] = p{-dx, dy}
			}
		case "U":
			if y == 0 {
				mv[w] = p{dx, -dy}
			}
		case "D":
			if y == 8 {
				mv[w] = p{dx, -dy}
			}
		case "RU":
			if x == 8 && y == 0 {
				mv[w] = p{-dx, -dy}
			} else if x == 8 {
				mv[w] = p{-dx, dy}
			} else if y == 0 {
				mv[w] = p{dx, -dy}
			}
		case "RD":
			if x == 8 && y == 8 {
				mv[w] = p{-dx, -dy}
			} else if x == 8 {
				mv[w] = p{-dx, dy}
			} else if y == 8 {
				mv[w] = p{dx, -dy}
			}
		case "LU":
			if x == 0 && y == 0 {
				mv[w] = p{-dx, -dy}
			} else if x == 0 {
				mv[w] = p{-dx, dy}
			} else if y == 0 {
				mv[w] = p{dx, -dy}
			}
		case "LD":
			if x == 0 && y == 8 {
				mv[w] = p{-dx, -dy}
			} else if x == 0 {
				mv[w] = p{-dx, dy}
			} else if y == 8 {
				mv[w] = p{dx, -dy}
			}
		}
		x += mv[w].dx
		y += mv[w].dy
	}
	wr.WriteString("\n")
	wr.Flush()
}
