package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var sx, sy, tx, ty int
	fmt.Scan(&sx, &sy, &tx, &ty)

	var wr = bufio.NewWriter(os.Stdout)
	dx, dy := tx-sx, ty-sy
	for i := 0; i < dy; i++ {
		wr.WriteString("U")
	}
	for i := 0; i < dx; i++ {
		wr.WriteString("R")
	}

	for i := 0; i < dy; i++ {
		wr.WriteString("D")
	}
	for i := 0; i < dx; i++ {
		wr.WriteString("L")
	}

	wr.WriteString("L")
	for i := 0; i < dy+1; i++ {
		wr.WriteString("U")
	}
	for i := 0; i < dx+1; i++ {
		wr.WriteString("R")
	}
	wr.WriteString("D")

	wr.WriteString("R")
	for i := 0; i < dy+1; i++ {
		wr.WriteString("D")
	}
	for i := 0; i < dx+1; i++ {
		wr.WriteString("L")
	}
	wr.WriteString("U")
	wr.WriteString("\n")
	wr.Flush()
}
