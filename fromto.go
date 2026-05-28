package main

import "fmt"

func FromTo(from, to int) string {
	if from > 99 || from < 0 || to > 99 || to < 0 {
		return "Invalid\n"
	}

	// FIX 1: Use 'res' everywhere so it matches the rest of your code
	res := ""
	step := 1
	if from > to {
		step = -1
	}

	for {
		// FIX 2 & 3: Simplified manual 2-digit conversions using math.
		// Since all numbers are confirmed 0-99, the tens digit is always from/10
		// and the units digit is always from%10. This works for ALL numbers 0-99.
		tens := from / 10
		units := from % 10

		res += string(rune('0'+tens)) + string(rune('0'+units))

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
