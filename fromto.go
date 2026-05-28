package main

import (
    "fmt"
    "strconv"
    )
func FromTo(from int, to int) string {
if (from <0 || from > 99) || (to < 0 || to > 99){
    return "Invalid\n"
}
step:= 1
result := ""
if from > to {
    step = -1
}

for {
    tens:= from /10
    unit := from % 10
    result += strconv.Itoa(tens) + strconv.Itoa(unit)

if from == to {
    break
}
result += ", "
from += step
}
return result + "\n"
}

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}