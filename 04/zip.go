package main

import  (

"fmt"
"strconv"
)

func ZipString(s string) string {
	if len(s) == 0 {
		return ""
	}

	count := 1
	result := ""

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && s[i] == s[i+1] {
			count++
		} else {
			result += strconv.Itoa(count) + string(s[i])
			count = 1
		}
	}
	return result
}
func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}
