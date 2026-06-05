package main

import (
	"fmt"
	"os"
)

func main() {
	// if arguments are not exactly 2, print nothing and stop
	if len(os.Args) != 3 {
		return
	}

	s1, s2 := os.Args[1], os.Args[2]

	// i tracks where we are in s1
	// j tracks where we are in s2
	i, j := 0, 0

	// keep going as long as we have not finished either string
	for i < len(s1) && j < len(s2) {
		// if the current characters match, move forward in s1
		if s1[i] == s2[j] {
			i++
		}
		// always move forward in s2
		j++
	}

	// if i reached the end of s1, we matched every character
	if i == len(s1) {
		fmt.Println(s1)
	}
}