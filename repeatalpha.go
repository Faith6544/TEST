// Online Go compiler to run Golang program online
// Print "Start small. Ship something." message

package main

import "fmt"

func RepeatAlpha(s string) string {
	result := ""
	for i, ch := range s {

		for j := 0; j < i+1; j++ {
			result += string(ch)
		}

	}
	return result
}
func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}
