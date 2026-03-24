package piscine

import (
	"os"
	"strconv"
	"fmt"
)

func mainQuadA() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: quadA x y")
		os.Exit(1)
	}
	x, _ := strconv.Atoi(os.Args[1])
	y, _ := strconv.Atoi(os.Args[2])
	QuadA(x, y)
}
