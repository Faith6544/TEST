package main

import "fmt"

func CamelToSnakeCase(s string) string {
	if len(s) == 0 {
		return ""
	}
	for i, c := range s {
		if c < 'a' && c > 'z' || c < 'A' && c > 'Z' {
			return s
		}
		if c >= 'A' && c <= 'Z' && i+1 < len(s) && s[i+1] >= 'A' && s[i+1] <= 'Z' {
			return s
		}
		if c >= 'A' && c <= 'Z' && i == len(s)-1 {
			return s
		}
	}
	result := ""
	for i, c := range s {
		if c >= 'A' && c <= 'Z' && i > 0 {
			result += "_"
		}
		if c >= 'A' && c <= 'Z' {
			result += string(c + 32)
		} else {
			result += string(c)
		}
	}
	return result

}
func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
