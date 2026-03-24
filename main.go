package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func makeLine(x int, leftEdge, filler, rightEdge rune) string {
	row := string(leftEdge)
	for i := 0; i < x-2; i++ {
		row += string(filler)
	}
	if x > 1 {
		row += string(rightEdge)
	}
	return row
}

func generateQuadA(x, y int) string {
	if x < 1 || y < 1 {
		return ""
	}
	var lines []string
	lines = append(lines, makeLine(x, 'o', '-', 'o'))
	for i := 0; i < y-2; i++ {
		lines = append(lines, makeLine(x, '|', ' ', '|'))
	}
	if y > 1 {
		lines = append(lines, makeLine(x, 'o', '-', 'o'))
	}
	return strings.Join(lines, "\n")
}

func generateQuadB(x, y int) string {
	if x < 1 || y < 1 {
		return ""
	}
	var lines []string
	lines = append(lines, makeLine(x, '/', '*', '\\'))
	for i := 0; i < y-2; i++ {
		lines = append(lines, makeLine(x, '*', ' ', '*'))
	}
	if y > 1 {
		lines = append(lines, makeLine(x, '\\', '*', '/'))
	}
	return strings.Join(lines, "\n")
}

func generateQuadC(x, y int) string {
	if x < 1 || y < 1 {
		return ""
	}
	var lines []string
	lines = append(lines, makeLine(x, 'A', 'B', 'A'))
	for i := 0; i < y-2; i++ {
		lines = append(lines, makeLine(x, 'B', ' ', 'B'))
	}

	if y > 1 {
		lines = append(lines, makeLine(x, 'C', 'B', 'C'))
	}
	return strings.Join(lines, "\n")
}

func generateQuadD(x, y int) string {
	if x < 1 || y < 1 {
		return ""
	}
	var lines []string
	lines = append(lines, makeLine(x, 'A', 'B', 'C'))
	for i := 0; i < y-2; i++ {
		lines = append(lines, makeLine(x, 'B', ' ', 'B'))
	}
	if y > 1 {
		lines = append(lines, makeLine(x, 'A', 'B', 'C'))
	}
	return strings.Join(lines, "\n")
}

func generateQuadE(x, y int) string {
	if x < 1 || y < 1 {
		return ""
	}
	var lines []string
	lines = append(lines, makeLine(x, 'A', 'B', 'C'))
	for i := 0; i < y-2; i++ {
		lines = append(lines, makeLine(x, 'B', ' ', 'B'))
	}
	if y > 1 {
		lines = append(lines, makeLine(x, 'C', 'B', 'A'))
	}
	return strings.Join(lines, "\n")
}

type QuadMatch struct {
	name string
	w, h int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if len(lines) == 0 {
		fmt.Println("Not a quad function")
		return
	}
	input := strings.Join(lines, "\n")
	h := len(lines)
	w := len(lines[0])
	for _, line := range lines {
		if len(line) != w {
			fmt.Println("Not a quad function")
			return
		}
	}
	generators := []struct {
		name string
		fn   func(int, int) string
	}{
		{"quadA", generateQuadA},
		{"quadB", generateQuadB},
		{"quadC", generateQuadC},
		{"quadD", generateQuadD},
		{"quadE", generateQuadE},
	}
	var matches []QuadMatch
	for _, g := range generators {
		if g.fn(w, h) == input {
			matches = append(matches, QuadMatch{g.name, w, h})
		}
	}
	if len(matches) == 0 {
		fmt.Println("Not a quad function")
		return
	}
	sort.Slice(matches, func(i, j int) bool {
		return matches[i].name < matches[j].name
	})
	parts := make([]string, len(matches))
	for i, m := range matches {
		parts[i] = fmt.Sprintf("[%s] [%d] [%d]", m.name, m.w, m.h)
	}
	fmt.Println(strings.Join(parts, " || "))
}
