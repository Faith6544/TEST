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

	for {
		// Format number with leading zero if less than 10
		if from < 10 {
			res += "0" + string(rune('0'+from))
		} else {
			res += string(rune('0'+from/10)) + string(rune('0'+from%10))
		}

		if from == to {
			break
		}
		res += ", "
		from += step
	}

	return res + "\n"
}

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}
