package main

import (
	"github.com/01-edu/z01"
)

func PrintMemory(arr [10]byte) {
	hex := "0123456789abcdef"
	for i, r := range arr {
		z01.PrintRune(rune(hex[r/16]))
		z01.PrintRune(rune(hex[r%16]))
		if i == 3 || i == 7 || i == 9 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}
	for _, r := range arr {
		if r < 32 || r > 126 {
			z01.PrintRune('.')
		} else {
			z01.PrintRune(rune(r))
		}
	}
	z01.PrintRune('\n')
}
