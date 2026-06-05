package main

import "fmt"

func ThirdTimeIsACharm(str string) string {

	if len(str) < 3 {
		return "\n"
	}
	output := ""
	for i := 2; i < len(str); i += 3 {
		output += string(str[i])
	}
	return output + "\n"
}
func main() {
	fmt.Print(ThirdTimeIsACharm("123456789"))
	fmt.Print(ThirdTimeIsACharm(""))
	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(ThirdTimeIsACharm("12"))
}
