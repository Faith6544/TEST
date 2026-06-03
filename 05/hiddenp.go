package main

import (
	"fmt"
	"os"
)

func HiddenP(s1, s2 string) {
	if s1 == "" {
		fmt.Println(1)
		return
	}

	j := 0
	for i := 0; i < len(s2); i++ {
		// if current character in s2 matches current character in s1
		if s2[i] == s1[j] {
			// move forward in s1
			j++
		}
		// if we have matched all characters in s1, print 1 and stop
		if j == len(s1) {
			fmt.Println(1)
			return
		}
	}
	// if we finished s2 without matching all of s1, print 0
	fmt.Println(0)
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	HiddenP(os.Args[1], os.Args[2])
}



package main

import (
    "os"
    "github.com/01-edu/z01"
)

func main() {
    if len(os.Args) != 3 {
        return
    }
    s1, s2 := os.Args[1], os.Args[2]
    i, j := 0, 0
    for i < len(s1) && j < len(s2) {
        if s1[i] == s2[j] {
            i++
        }
        j++
    }
    if i == len(s1) {
        z01.PrintRune('1')
    } else {
        z01.PrintRune('0')
    }
    z01.PrintRune('\n')
}