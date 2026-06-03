package main

import (
"strconv"
"fmt"
)

func FromTo(from int, to int) string {
	if from < 0 || to > 99 || to < 0 || from > 99 {
		return "Invalid\n"
	}
	step := 1
	var result string
	if from > to {
		step = -1
	}
	for i := from ; i != to +step ; i+=step{
		if i < 10 {
			result += "0" + strconv.Itoa(i)
		}else {
			result += strconv.Itoa(i)
		}
		if i != to {
			result += ", "
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
