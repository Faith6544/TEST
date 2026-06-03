package main

import "fmt"

func FishAndChips(n int) string {
	if n%2 == 0 && n%3 == 0 {
		return "Fish And Chips"
	}
	if n%2 == 0 {
		return "Fish"
	}
	if n%3 == 0 {
		return "Chips"
	}
	if n < 0 {
		return "error: number is negative"
	}
	return "non divisible"
}
func main() {
	fmt.Println(FishAndChips(4))
	fmt.Println(FishAndChips(9))
	fmt.Println(FishAndChips(6))
}
