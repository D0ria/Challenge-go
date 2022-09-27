package main

import (
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	piscine.ForEach(PrintNbr, a)
}

func PrintNbr(n int) {
	var tabint []int
	for i := n; i > 0; i /= 10 {
		tabint = append(tabint, i%10)
	}
	for i := len(tabint) - 1; i >= 0; i-- {
		z01.PrintRune(rune(tabint[i] + 48))
	}
}
