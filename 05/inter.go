package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	
	var seen [256]bool

	
	var alreadyPrinted [256]bool

	for i := 0; i < len(s2); i++ {
		seen[s2[i]] = true
	}

	
	for i := 0; i < len(s1); i++ {
		ch := s1[i]


		if seen[ch] && !alreadyPrinted[ch] {
			alreadyPrinted[ch] = true
			// print the character
			fmt.Printf("%c", ch)
		}
	}

	
	fmt.Println()
}