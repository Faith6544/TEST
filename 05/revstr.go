package main

import (
	"fmt"
	"os"
)

func main() {
	// go through each argument one by one
	for _, arg := range os.Args[1:] {
		// go through each character in the argument
		for i := 0; i < len(arg); i++ {
			c := arg[i]

			// check if this character is the last of a word
			// that means the next character is a space OR we are at the end
			if i+1 == len(arg) || arg[i+1] == ' ' {
				// this is the last letter of a word
				// if it is lowercase, make it uppercase
				if c >= 'a' && c <= 'z' {
					c -= 32
				}
			} else {
				// this is NOT the last letter of a word
				// if it is uppercase, make it lowercase
				if c >= 'A' && c <= 'Z' {
					c += 32
				}
			}

			// print the character
			fmt.Printf("%c", c)
		}
		// print a new line after each argument
		fmt.Println()
	}
}