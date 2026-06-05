package main

import "fmt"

func CanJump(slice []uint) bool {
	if len(slice) == 0 {
		return false
	}
	if len(slice) == 1 {
		return true
	}
	for i := 0; i < len(slice); i += int(slice[i]) {
		if i == len(slice)-1 {
			return true
		}
		if slice[i] == 0 {
			return false
		}
	}
	return false
}

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))
}
