package main

import "fmt"

func QuadA(x, y int) {
	if x < 1 || y < 1 {
		return
	}
	
	printHorizontalLine := func(leftEdge, filler, rightEdge rune) {
		fmt.Printf("%c", leftEdge)
		for i := 0; i < (x - 2); i++ {
			fmt.Printf("%c", filler)
		}
		if x > 1 {
			fmt.Printf("%c", rightEdge)
		}
		fmt.Println()
	}


	printHorizontalLine('o', '-', 'o')
	for i := 0; i < y-2; i++ {
		printHorizontalLine('|', ' ', '|')
	}
	if y > 1 {
		printHorizontalLine('o', '-', 'o')
	}
}
