package main

import "fmt"

func HashCode(dec string) string {
	hash := ""
	size := len(dec)
	for _, ch := range dec {
		c := (int(ch) + size) % 127

		// FIXED: Added '|| c == 127' to catch the unprintable DEL character
		if c < 32 || c == 127 {
			c += 33
		}
		hash += string(rune(c))
	}
	return hash
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
