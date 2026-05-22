package main

import "fmt"

func WeAreUnique(str1, str2 string) int {
	if len(str1) == 0 && len(str2) == 0 {
		return -1
	}
	count := 0

	for _, c := range str1 {
		found := false
		for _, d := range str2 {
			if c == d {
				found = true

			}
		}
		if !found {
			count++
		}
	}

	for _, c := range str2 {
		found := false
		for _, d := range str1 {
			if c == d {
				found = true

			}
		}
		if !found {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}
