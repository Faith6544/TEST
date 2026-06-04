package main

import (
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		for i := 0; i < len(arg); i++ {
			c := arg[i]
			if i+1 == len(arg) || arg[i+1] == ' ' {
				if c >= 'a' && c <= 'z' {
					c -= 32
				}
			} else {
				if c >= 'A' && c <= 'Z' {
					c += 32
				}
			}
			fmt.Print(c)

		}
		fmt.Print("\n")
	}
}
