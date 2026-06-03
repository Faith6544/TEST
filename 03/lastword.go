package main

import "fmt"

func LastWord(s string) string{
 i:= len(s)-1

 for i>=0 && s[i]==' '{
	i--
 }

 end := i+1
  for i>=0 && s[i] !=' '{
	i--
 }
return s[i+1 : end] + "\n"
}
func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}