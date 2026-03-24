package piscine

import (
	"os"
	"strconv"
	"fmt"
)

func mainQuadE() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: quadE x y")
		os.Exit(1)
	}
	x, _ := strconv.Atoi(os.Args[1])
	y, _ := strconv.Atoi(os.Args[2])
	QuadE(x, y)
}
