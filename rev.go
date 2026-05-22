package main

import "fmt"

func revcomb(n int) int {
	count := 0

	for i := 9; i >= 2; i-- {
		for j := i - 1; j >= 1; j-- {
			for k := j - 1; k >= 0; k-- {
				fmt.Printf("%d%d%d", i, j, k)
				count++
				if i == 2 && j == 1 && k == 0 {
					fmt.Print("\n")
				} else {
					fmt.Print(", ")

				}
			}
		}
	}

	return count
}

func main() {
	total := revcomb(3)
	fmt.Printf("Total combinations printed: %d\n", total)
}
