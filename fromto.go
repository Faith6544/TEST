package main

import "fmt"

func FromTo(from, to int) string {
	if from > 99 || from < 0 || to > 99 || to < 0 {
		return "Invalid\n"
	}

	result := ""
	step := 1
	if from > to {
		step = -1
	}

	for i := from; i != to+step; i += step {
		if result != "" {
			result += ", "
		}
		if i < 10 {
			result += "0" + string(rune('0'+i))
		} else {
			result += string(rune('0'+i/10)) + string(rune('0'+i%10))
		}
	}
	return result + "\n"
}

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}
