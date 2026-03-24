package main

import "fmt"

func QuadC(x, y int) {
	if x < 1 || y < 1 {
		return
	}

	printHorizontalLine := func(x int, leftEdge, filler, rightEdge rune) {
		fmt.Printf("%c", leftEdge)
		for i := 0; i < (x - 2); i++ {
			fmt.Printf("%c", filler)
		}
		if x > 1 {
			fmt.Printf("%c", rightEdge)
		}
		fmt.Println()
	}

	printHorizontalLine(x, 'A', 'B', 'A')

	for i := 0; i < y-2; i++ {
		printHorizontalLine(x, 'B', ' ', 'B')
	}

	if y > 1 {
		printHorizontalLine(x,'C', 'B', 'C')
	}
}