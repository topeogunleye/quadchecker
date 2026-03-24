package piscine

import (
	"os"
	"strconv"
	"fmt"
)

func mainQuadD() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: quadD x y")
		os.Exit(1)
	}
	x, _ := strconv.Atoi(os.Args[1])
	y, _ := strconv.Atoi(os.Args[2])
	QuadD(x, y)
}
