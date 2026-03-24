package main

import "github.com/01-edu/z01"

func QuadD(x, y int) {
	if x < 1 || y < 1 {
		return
	}

	printHorizontalLine := func(x int, leftEdge, filler, rightEdge int32) {
		z01.PrintRune(leftEdge)
		for i := 0; i < (x - 2); i++ {
			z01.PrintRune(filler)
		}

		if x > 1 {
			z01.PrintRune(rightEdge)
		}
		z01.PrintRune('\n')
	}

	printHorizontalLine(x, 'A', 'B', 'C')

	for i := 0; i < y-2; i++ {
		printHorizontalLine(x, 'B', ' ', 'B')
	}

	if y > 1 {
		printHorizontalLine(x, 'A', 'B', 'C')
	}
}
