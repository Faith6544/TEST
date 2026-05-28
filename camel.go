package main

import "fmt"

func CamelToSnakeCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	for i, c := range s {
		// FIX 1: Corrected the character filter.
		// If it's NOT a lowercase letter AND it's NOT an uppercase letter, it's invalid.
		if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') {
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
		// FIX 2: Removed the '+ 32' conversion because the expected
		// output requires keeping the original uppercase letters intact.
		result += string(c)
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
