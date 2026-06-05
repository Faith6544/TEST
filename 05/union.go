package main

import (
    "os"
    "github.com/01-edu/z01"
)

func main() {
    if len(os.Args) == 3 {
        var seen [256]bool
        for _, s := range os.Args[1:] {
            for i := 0; i < len(s); i++ {
                if !seen[s[i]] {
                    seen[s[i]] = true
                    z01.PrintRune(rune(s[i]))
                }
            }
        }
    }
    z01.PrintRune('\n')
}