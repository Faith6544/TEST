package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

// prints a number digit by digit
func printInt(n int) {
	if n >= 10 {
		printInt(n / 10) // print left digits first
	}
	z01.PrintRune(rune('0' + n%10)) // print last digit
}

// checks if a number is prime
func Prime(n int) bool {
	if n < 2 {
		return false // 0 and 1 are not prime
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false // divisible by something, not prime
		}
	}
	return true // nothing divided it, so it is prime
}

func main() {
	// need exactly one argument
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	// convert argument from string to number
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		z01.PrintRune('\n')
		return
	}

	// add up all prime numbers from 2 to n
	sum := 0
	for i := 2; i <= n; i++ {
		if Prime(i) {
			sum += i
		}
	}

	// print the result
	printInt(sum)
	z01.PrintRune('\n')
}
