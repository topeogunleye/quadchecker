package piscine

import "github.com/01-edu/z01"

func QuadC(x, y int) {
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

	if y >= 1 {
		printHorizontalLine(x, 'A', 'B', 'A')
	}

	if y > 2 {
		for i := 0; i < y-2; i++ {
			printHorizontalLine(x, 'B', ' ', 'B')
		}
	}

	if y > 1 {
		printHorizontalLine(x, 'C', 'B', 'C')
	}
}
