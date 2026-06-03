package main

import (
	"fmt"
	"strings"
)

func WeAreUnique(str1, str2 string) int {
	if len(str1) == 0 && len(str2) == 0 {
		return -1
	}
	count := 0
	unique := make(map[rune]bool)
	for _, c := range str1 {
		if !strings.Contains(str2, string(c)) && !unique[c] {
			unique[c] = true
			count++
		}
	}
	for _, c := range str2 {
		if !strings.Contains(str1, string(c)) && !unique[c] {
			unique[c] = true
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
