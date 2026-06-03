// Online Go compiler to run Golang program online
// Print "Start small. Ship something." message

package main
import "fmt"

func RepeatAlpha(s string) string {
<<<<<<< HEAD
	result := ""
	count := 1
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			count = int(ch - 'a' + 1)
			for i := 0; i < count; i++ {
				result += string(ch)
			}

		} else if ch >= 'A' && ch <= 'Z' {
			count = int(ch - 'A' + 1)
			for i := 0; i < count; i++ {
				result += string(ch)
			}

		} else {
			result += string(ch)
		}
	}
	return result

=======
    res:= ""
    for _, ch := range s {
        if ch >='a' && ch <= 'z' {
            count := int( ch -'a' + 1)
            for i:= 0 ; i < count ; i++{
                res +=string(ch)
            }
        } else if ch >='A' && ch <= 'Z' {
            count := int( ch -'A' + 1)
            for i:= 0 ; i < count ; i++{
                res +=string(ch)
            }
        } else {
            res+= string(ch)
        }
    }
    return res
>>>>>>> 6fb42b27e7b27f58d0f98d87b5c2244554294b4b
}
 func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}