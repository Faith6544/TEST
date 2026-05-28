package main

import "fmt"

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	result:=""

	negative := n <0 
	if negative {
		n=-n
	}
	for n> 0  {
		result = string(byte('0' + n %10)) + result
		n /=10

	}
	if negative{
		result= "-" + result
	}
	return result
}

func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}
