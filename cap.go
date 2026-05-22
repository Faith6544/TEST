package main

import "github.com/01-edu/z01"

// Notice: No return type here! This matches "want ()"
func PrintMemory(arr [10]byte) {
	hexDigits := "0123456789abcdef"
	output := ""

	// 1. Build the Hex strings
	for i, b := range arr {
		output += string(hexDigits[b/16]) + string(hexDigits[b%16])
		if i == 3 || i == 7 || i == 9 {
			output += "\n"
		} else {
			output += " "
		}
	}

	// 2. Build the ASCII characters
	for _, b := range arr {
		if b >= 32 && b <= 126 {
			output += string(b)
		} else {
			output += "."
		}
	}
	output += "\n"

	// 3. Print the final combined text block letter by letter
	for _, ch := range output {
		z01.PrintRune(ch)
	}
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
